package mysql

import (
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"task-management/internal/types"
	"testing"
)

func TestAuthMySQL_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error `%s` was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthMySQL(db)

	tests := []struct {
		name    string
		mock    func()
		input   types.User
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectExec("INSERT INTO users").
					WithArgs("test", "test", "password").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery("SELECT LAST_INSERT_ID()").
					WillReturnRows(rows)
			},
			input: types.User{
				Name:     "test",
				Username: "test",
				Password: "password",
			},
			want: 1,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.CreateUser(testCase.input)
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
