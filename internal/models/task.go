package models

import "time"

type Task struct {
	TaskID      int       `json:"task_id"`
	UserID      int       `json:"user_id"`
	TaskName    string    `json:"task_name"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
}
