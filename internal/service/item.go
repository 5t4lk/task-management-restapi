package service

import (
	"task-management/internal/repository"
	"task-management/internal/types"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (i *ItemService) Create(userId, taskId int, item types.TaskItem) (int, error) {
	return i.repo.Create(taskId, item)
}
