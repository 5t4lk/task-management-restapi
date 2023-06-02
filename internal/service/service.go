package service

import (
	"task-management/internal/repository"
	"task-management/internal/types"
)

type Authorization interface {
	CreateUser(user types.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
