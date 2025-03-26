package main

import (
	"log"
	"project/internal/database"
	"project/internal/handlers"
	"project/internal/taskService"
	"project/internal/userService"
	"project/internal/web/tasks"
	"project/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()

	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	if err := database.DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	TaskRepo := taskService.NewTaskRepository(database.DB)
	TaskService := taskService.NewTaskService(TaskRepo)
	TaskHandler := handlers.NewTaskHandler(TaskService)

	UserRepo := userService.NewUserRepository(database.DB)
	UserService := userService.NewUserService(UserRepo)
	UserHandler := handlers.NewUserHandler(UserService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTaskHandler := tasks.NewStrictHandler(TaskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(UserHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
