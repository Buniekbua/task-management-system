package handlers

import (
	"net/http"
	"strconv"

	"github.com/buniekbua/task-managment-system/internal/models"
	"github.com/buniekbua/task-managment-system/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

func (h *TaskHandler) CreateTaskHandler(c echo.Context) error {
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := h.taskUseCase.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error creating task"})
	}
	return c.JSON(http.StatusCreated, task)

}

func (h *TaskHandler) GetTaskByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task id"})
	}

	task, err := h.taskUseCase.GetTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTaskHandler(c echo.Context) error {
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task id"})
	}

	if err := h.taskUseCase.UpdateTask(id, &task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error updating task"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTaskHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task id"})
	}

	if err := h.taskUseCase.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error deleting task"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Task deleted successfully"})
}

func (h *TaskHandler) GetTasksByUserIDHandler(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user id"})
	}

	tasks, err := h.taskUseCase.GetTasksByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error getting tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}
