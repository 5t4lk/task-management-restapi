package service

import (
	"task-management/internal/repository"
	"task-management/internal/types"
)

type ItemService struct {
	repo     repository.Item
	taskRepo repository.Task
}

func NewItemService(repo repository.Item, taskRepo repository.Task) *ItemService {
	return &ItemService{repo: repo, taskRepo: taskRepo}
}

func (i *ItemService) Create(userId, taskId int, item types.TaskItem) (int, error) {
	_, err := i.taskRepo.GetById(userId, taskId)
	if err != nil {
		return 0, err
	}

	return i.repo.Create(taskId, item)
}

func (i *ItemService) GetAll(userId, taskId int) ([]types.TaskItem, error) {
	return i.repo.GetAll(userId, taskId)
}

func (i *ItemService) GetById(userId, itemId int) (types.TaskItem, error) {
	return i.repo.GetById(userId, itemId)
}

func (i *ItemService) Update(userId, itemId int, input types.UpdateItemInput) error {
	return i.repo.Update(userId, itemId, input)
}
