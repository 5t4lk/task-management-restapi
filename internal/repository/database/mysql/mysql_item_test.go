package mysql

import (
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"task-management/internal/types"
	"testing"
)

func TestNewItemMySQL(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connector", err)
	}
	defer db.Close()

	r := NewItemMySQL(db)

	type args struct {
		taskId int
		item   types.TaskItem
	}
	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				taskId: 1,
				item: types.TaskItem{
					Title:       "test title",
					Description: "test description",
				},
			},
			mock: func(args args, id int) {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO items").
					WithArgs("test title", "test description").
					WillReturnResult(sqlmock.NewResult(1, 1))

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("SELECT LAST_INSERT_ID()").
					WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO tasks_items").
					WithArgs(args.taskId, id).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			want: 1,
		},
		{
			name: "Empty fields",
			mock: func(args args, id int) {
				mock.ExpectBegin()

				mock.ExpectExec("INSERT INTO tasks").
					WithArgs("", "test description", "test status", "test end date").
					WillReturnResult(sqlmock.NewResult(0, 0))

				mock.ExpectRollback()
			},
			input: args{
				taskId: 1,
				item: types.TaskItem{
					Title:       "",
					Description: "test description",
				},
			},
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.input, testCase.want)

			got, err := r.Create(testCase.input.taskId, testCase.input.item)
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
