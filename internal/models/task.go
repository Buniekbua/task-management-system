package models

import "time"

type Task struct {
	TaskID      int       `db:"task_id" json:"task_id"`
	UserID      int       `db:"user_id" json:"user_id"`
	TaskName    string    `db:"task_name" json:"task_name"`
	Description string    `db:"description" json:"description"`
	DueDate     string    `db:"due_date" json:"due_date"`
	Priority    string    `db:"priority" json:"priority"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
