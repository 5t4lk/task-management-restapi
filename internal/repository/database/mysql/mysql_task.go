package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"task-management/internal/types"
)

type TaskMySQL struct {
	db *sqlx.DB
}

func NewTaskMySQL(db *sqlx.DB) *TaskMySQL {
	return &TaskMySQL{db: db}
}

func (t *TaskMySQL) Create(userId int, task types.Task) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description, status, end_date) VALUES (?, ?, ?, ?)", taskTable)
	result, err := tx.Exec(createTaskQuery, task.Title, task.Description, task.Status, task.EndDate)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersTaskQuery := fmt.Sprintf("INSERT INTO %s (user_id, task_id) VALUES (?, ?)", userTasksTable)
	_, err = tx.Exec(createUsersTaskQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}

func (t *TaskMySQL) GetAll(userId int) ([]types.Task, error) {
	var tasks []types.Task

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.status, tl.end_date FROM %s tl INNER JOIN %s ul on tl.id = ul.task_id WHERE ul.user_id = ?",
		taskTable, userTasksTable)
	err := t.db.Select(&tasks, query, userId)

	return tasks, err
}

func (t *TaskMySQL) GetById(userId int, taskId int) (types.Task, error) {
	var task types.Task

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.status, tl.end_date FROM %s tl INNER JOIN %s ul on tl.id = ul.task_id WHERE ul.user_id = ? AND ul.task_id = ?",
		taskTable, userTasksTable)
	err := t.db.Get(&task, query, userId, taskId)

	return task, err
}

func (t *TaskMySQL) Update(userId, listId int, input types.UpdateTaskInput) error {
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

	if input.EndDate != nil {
		setValues = append(setValues, fmt.Sprintf("end_date = ?"))
		args = append(args, *input.EndDate)
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status = ?"))
		args = append(args, *input.Status)
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = ?"))
		args = append(args, *input.Done)
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl INNER JOIN %s ul ON tl.id = ul.task_id SET %s WHERE ul.task_id = ? AND ul.user_id = ?",
		taskTable, userTasksTable, setQuery)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %v", args)

	_, err := t.db.Exec(query, args...)
	return err
}
