package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"task-management/internal/types"
)

type AuthMySQL struct {
	db *sqlx.DB
}

func NewAuthMySQL(db *sqlx.DB) *AuthMySQL {
	return &AuthMySQL{db: db}
}

func (a *AuthMySQL) CreateUser(user types.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values (?, ?, ?)", usersTable)

	result, err := a.db.Exec(query, user.Name, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (a *AuthMySQL) GetUser(username, password string) (types.User, error) {
	var user types.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = ? AND password_hash = ?", usersTable)
	err := a.db.Get(&user, query, username, password)

	return user, err
}
