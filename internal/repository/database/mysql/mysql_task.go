package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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
