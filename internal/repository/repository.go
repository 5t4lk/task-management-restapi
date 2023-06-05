package repository

import (
	"github.com/jmoiron/sqlx"
	"task-management/internal/repository/database/mysql"
	"task-management/internal/types"
)

type Authorization interface {
	CreateUser(user types.User) (int, error)
	GetUser(username, password string) (types.User, error)
}

type Task interface {
	Create(userId int, task types.Task) (int, error)
	GetAll(userId int) ([]types.Task, error)
	GetById(userId int, taskId int) (types.Task, error)
	Update(userId, taskId int, input types.UpdateTaskInput) error
	Delete(userId, taskId int) error
}

type Item interface {
	Create(taskId int, item types.TaskItem) (int, error)
	GetAll(userId, taskId int) ([]types.TaskItem, error)
	GetById(userId, itemId int) (types.TaskItem, error)
	Update(userId, itemId int, input types.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Repository struct {
	Authorization
	Task
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: mysql.NewAuthMySQL(db),
		Task:          mysql.NewTaskMySQL(db),
		Item:          mysql.NewItemMySQL(db),
	}
}
