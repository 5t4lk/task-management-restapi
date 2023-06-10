package service

import (
	"task-management/internal/repository"
	"task-management/internal/types"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user types.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Task interface {
	Create(userId int, task types.Task) (int, error)
	GetAll(userId int) ([]types.Task, error)
	GetById(userId int, taskId int) (types.Task, error)
	Update(userId, taskId int, input types.UpdateTaskInput) error
	Delete(userId, taskId int) error
}

type Item interface {
	Create(userId, taskId int, item types.TaskItem) (int, error)
	GetAll(userId, taskId int) ([]types.TaskItem, error)
	GetById(userId, itemId int) (types.TaskItem, error)
	Update(userId, itemId int, input types.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Authorization
	Task
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Task:          NewTaskService(repos.Task),
		Item:          NewItemService(repos.Item, repos.Task),
	}
}
