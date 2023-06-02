package mysql

import (
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

}
