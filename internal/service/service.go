package service

import (
	"task-management/internal/repository"
	"task-management/internal/types"
)

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

type Service struct {
	Authorization
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Task:          NewTaskService(repos.Task),
	}
}
