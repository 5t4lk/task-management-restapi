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
		{
			name: "Empty fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectExec("INSERT INTO users").
					WithArgs("test", "", "password").
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectQuery("SELECT LAST_INSERT_ID()").
					WillReturnRows(rows)
			},
			input: types.User{
				Name:     "test",
				Username: "",
				Password: "password",
			},
			wantErr: true,
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

func TestAuthMySQL_GetUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error `%s` was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthMySQL(db)

	type args struct {
		username string
		password string
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    types.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password"}).
					AddRow(1, "test", "test", "password")
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("test", "password").WillReturnRows(rows)
			},
			input: args{
				username: "test",
				password: "password",
			},
			want: types.User{
				Id:       1,
				Name:     "test",
				Username: "test",
				Password: "password",
			},
		},
		{
			name: "Not found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password"})
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("test", "password").
					WillReturnRows(rows)
			},
			input: args{
				username: "test",
				password: "password",
			},
			wantErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.GetUser(testCase.input.username, testCase.input.password)
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
