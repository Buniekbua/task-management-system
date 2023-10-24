package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/buniekbua/task-managment-system/internal/models"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

type DBError struct {
	Err error
	Ctx string
}

func (e *DBError) Error() string {
	return fmt.Sprintf("%s: %v", e.Ctx, e.Err)
}

// Create task in database
func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := `
	INSERT INTO tasks (user_id, task_name, description, due_date, priority, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING task_id
	`
	err := r.db.QueryRow(query, task.UserID, task.TaskName, task.Description, task.DueDate, task.Priority, time.Now()).Scan(&task.TaskID)
	if err != nil {
		return &DBError{
			Err: err,
			Ctx: "Failed to create task",
		}
	}

	return nil
}

// Get task by ID from database
func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	query := `SELECT * FROM tasks WHERE task_id = $1`

	var task models.Task
	err := r.db.Get(&task, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, &DBError{
			Err: err,
			Ctx: "Failed to get task by ID",
		}
	}

	return &task, nil
}

// Updates task
func (r *TaskRepository) UpdateTask(id int, task *models.Task) error {
	query := `
		UPDATE tasks
		SET task_name = $2, description = $3, due_date = $4, priority = $5
		WHERE task_id = $1
	`
	_, err := r.db.Exec(query, id, task.TaskName, task.Description, task.DueDate, task.Priority)
	if err != nil {
		return &DBError{
			Err: err,
			Ctx: "Failed to update task",
		}
	}

	return nil
}

// Deletes task
func (r *TaskRepository) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE task_id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return &DBError{
			Err: err,
			Ctx: "Failed to delete task",
		}
	}

	return nil
}

// List all tasks for specific user
func (r *TaskRepository) GetTasksByUserID(userID int) ([]*models.Task, error) {
	query := `SELECT * FROM tasks WHERE user_id = $1`

	var tasks []*models.Task
	err := r.db.Select(&tasks, query, userID)
	if err != nil {
		return nil, &DBError{
			Err: err,
			Ctx: "Failed to get tasks by user ID",
		}
	}

	return tasks, nil
}
