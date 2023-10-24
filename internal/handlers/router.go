package handlers

import (
	"github.com/buniekbua/task-managment-system/internal/usecase"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, taskUseCase usecase.TaskUseCase, userUseCase usecase.UserUseCase) {
	taskHandler := NewTaskHandler(taskUseCase)
	userHandler := NewUserHandler(userUseCase)

	e.POST("/tasks", taskHandler.CreateTaskHandler)
	e.GET("/tasks/:id", taskHandler.GetTaskByIDHandler)
	e.PUT("/tasks/:id", taskHandler.UpdateTaskHandler)
	e.DELETE("/tasks/:id", taskHandler.DeleteTaskHandler)
	e.GET("/users/:id/tasks", taskHandler.GetTasksByUserIDHandler)

	e.POST("/users", userHandler.CreateUserHandler)
	e.GET("/users/:id", userHandler.GetUserByIDHandler)
	e.PUT("/users/:id", userHandler.UpdateUserHandler)
	e.DELETE("/users/:id", userHandler.DeleteUserHandler)

}
