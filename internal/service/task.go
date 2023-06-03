package service

import (
	"task-management/internal/repository"
	"task-management/internal/types"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) Create(userId int, task types.Task) (int, error) {
	return t.repo.Create(userId, task)
}
