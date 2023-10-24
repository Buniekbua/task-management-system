package usecase

import (
	"github.com/buniekbua/task-managment-system/internal/db"
	"github.com/buniekbua/task-managment-system/internal/models"
)

type TaskUseCase interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id int) (*models.Task, error)
	UpdateTask(id int, task *models.Task) error
	DeleteTask(id int) error
	GetTasksByUserID(userID int) ([]*models.Task, error)
}

type TaskUseCaseImpl struct {
	taskRepository db.TaskRepository
}

func NewTaskUseCase(tr db.TaskRepository) *TaskUseCaseImpl {
	return &TaskUseCaseImpl{taskRepository: tr}
}

func (uc *TaskUseCaseImpl) CreateTask(task *models.Task) error {
	return uc.taskRepository.CreateTask(task)
}

func (uc *TaskUseCaseImpl) GetTaskByID(id int) (*models.Task, error) {
	return uc.taskRepository.GetTaskByID(id)
}

func (uc *TaskUseCaseImpl) UpdateTask(id int, task *models.Task) error {
	return uc.taskRepository.UpdateTask(id, task)
}

func (uc *TaskUseCaseImpl) DeleteTask(id int) error {
	return uc.taskRepository.DeleteTask(id)
}

func (uc *TaskUseCaseImpl) GetTasksByUserID(userID int) ([]*models.Task, error) {
	return uc.taskRepository.GetTasksByUserID(userID)
}
