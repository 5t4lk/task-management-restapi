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

	createUsersTaskQuery := fmt.Sprintf("INSERT INTO %s (user_id, task_id) VALUES (?, ?)", userTasks)
	_, err = tx.Exec(createUsersTaskQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}
