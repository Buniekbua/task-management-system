package main

import (
	"github.com/buniekbua/task-managment-system/internal/db"
	"github.com/buniekbua/task-managment-system/internal/handlers"
	"github.com/buniekbua/task-managment-system/internal/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	dbConn, err := sqlx.Connect("postgres", "postgresql://root:password@localhost:5432/task_managment_db?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	taskRepo := db.NewTaskRepository(dbConn)
	taskUseCase := usecase.NewTaskUseCase(*taskRepo)

	userRepo := db.NewUserRepository(dbConn)
	userUseCase := usecase.NewUserUseCase(*userRepo)

	e := echo.New()
	handlers.NewRouter(e, taskUseCase, userUseCase)

	e.Logger.Fatal(e.Start(":8080"))
}
