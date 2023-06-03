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

func (t *TaskService) GetAll(userId int) ([]types.Task, error) {
	return t.repo.GetAll(userId)
}

func (t *TaskService) GetById(userId int, taskId int) (types.Task, error) {
	return t.repo.GetById(userId, taskId)
}
