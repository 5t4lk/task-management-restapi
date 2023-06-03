package types

type Task struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description" binding:"required"`
	Status      string `json:"status" db:"status"`
	EndDate     string `json:"endDate" db:"end_date"`
}

type UsersTask struct {
	Id     int
	UserId int
	TaskId int
}

type TaskItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Status      string `json:"status" db:"status"`
	EndDate     string `json:"endDate" db:"end_date"`
	Done        bool   `json:"done" db:"done"`
}

type TasksItem struct {
	Id     int
	TaskId int
	ItemId int
}
