package types

type Task struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
	Status      string `json:"status" db:"status"`
	EndDate     string `json:"end_date" db:"end_date"`
}

type UsersTask struct {
	Id     int `json:"id" db:"id"`
	UserId int `json:"userId" db:"user_id"`
	TaskId int `json:"taskId" db:"task_id"`
}

type TaskItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type TasksItem struct {
	Id     int
	TaskId int
	ItemId int
}

type GetAllTasksResponse struct {
	Data []Task `json:"data"`
}

type UpdateTaskInput struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Status      *string `json:"status" db:"status"`
	EndDate     *string `json:"end_date" db:"end_date"`
	Done        *bool   `json:"done" db:"done"`
}
