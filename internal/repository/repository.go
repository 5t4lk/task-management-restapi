package repository

import (
	"github.com/jmoiron/sqlx"
	"task-management/internal/repository/database/mysql"
	"task-management/internal/types"
)

type Authorization interface {
	CreateUser(user types.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: mysql.NewAuthMySQL(db),
	}
}
