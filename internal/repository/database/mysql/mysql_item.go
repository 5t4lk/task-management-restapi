package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"task-management/internal/types"
)

type itemMySQL struct {
	db *sqlx.DB
}

func NewItemMySQL(db *sqlx.DB) *itemMySQL {
	return &itemMySQL{db: db}
}

func (i *itemMySQL) Create(taskId int, item types.TaskItem) (int, error) {
	tx, err := i.db.Begin()
	if err != nil {
		return 0, err
	}

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", itemsTable)
	result, err := tx.Exec(createItemQuery, item.Title, item.Description)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	itemId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createTaskItemsQuery := fmt.Sprintf("INSERT INTO %s (task_id, item_id) VALUES (?, ?)", tasksItemsTable)
	_, err = tx.Exec(createTaskItemsQuery, taskId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(itemId), nil
}

func (i *itemMySQL) GetAll(userId, taskId int) ([]types.TaskItem, error) {
	var items []types.TaskItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.task_id = li.task_id WHERE li.task_id = ? AND ul.user_id = ?",
		itemsTable, tasksItemsTable, userTasksTable)
	if err := i.db.Select(&items, query, taskId, userId); err != nil {
		return nil, err
	}

	return items, nil
}
