package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (i *itemMySQL) GetById(userId, itemId int) (types.TaskItem, error) {
	var item types.TaskItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.task_id = li.task_id WHERE ti.id = ? AND ul.user_id = ?",
		itemsTable, tasksItemsTable, userTasksTable)

	if err := i.db.Get(&item, query, itemId, userId); err != nil {
		return types.TaskItem{}, err
	}

	return item, nil
}

func (i *itemMySQL) Update(userId, itemId int, input types.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = ?"))
		args = append(args, *input.Title)
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = ?"))
		args = append(args, *input.Description)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ti JOIN %s li ON ti.id = li.item_id JOIN %s ul ON li.task_id = ul.task_id SET %s WHERE ul.user_id = ? AND ti.id = ?",
		itemsTable, tasksItemsTable, userTasksTable, setQuery)
	args = append(args, userId, itemId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := i.db.Exec(query, args...)

	return err
}

func (i *itemMySQL) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE ti FROM %s ti
			JOIN %s li ON ti.id = li.item_id
			JOIN %s ul ON li.task_id = ul.task_id
			WHERE ul.user_id = ? AND ti.id = ?`,
		itemsTable, tasksItemsTable, userTasksTable)
	_, err := i.db.Exec(query, userId, itemId)

	return err
}
