package service

import (
	"task-management/internal/repository"
	"task-management/internal/types"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user types.User) (int, error) {
	return a.repo.CreateUser(user)
}
