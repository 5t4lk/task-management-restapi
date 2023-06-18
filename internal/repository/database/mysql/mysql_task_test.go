package mysql

import (
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"task-management/internal/types"
	"testing"
)

func TestItemMySQL_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connector", err)
	}
	defer db.Close()

	r := NewTaskMySQL(db)

	type args struct {
		userId int
		task   types.Task
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO tasks").
					WithArgs("test title", "test description", "test status", "test end date").
					WillReturnResult(sqlmock.NewResult(1, 1))

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("SELECT LAST_INSERT_ID()").
					WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO users_tasks").
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			input: args{
				userId: 1,
				task: types.Task{
					Title:       "test title",
					Description: "test description",
					Status:      "test status",
					EndDate:     "test end date",
				},
			},
			want: 1,
		},
		{
			name: "Empty fields",
			mock: func() {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO tasks").
					WithArgs("", "test description", "test status", "test end date").
					WillReturnResult(sqlmock.NewResult(0, 0))

				mock.ExpectRollback()
			},
			input: args{
				userId: 1,
				task: types.Task{
					Title:       "",
					Description: "test description",
					Status:      "test status",
					EndDate:     "test end date",
				},
			},
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.Create(testCase.input.userId, testCase.input.task)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
