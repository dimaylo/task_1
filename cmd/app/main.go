package main

import (
	"REST_API/internal/database"
	"REST_API/internal/handlers"
	"REST_API/internal/taskService"
	"REST_API/internal/userService"
	"REST_API/internal/web/tasks"
	"REST_API/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewTaskService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	userRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewUserService(userRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, strictTasksHandler)

	strictUsersHandler := users.NewStrictHandler(usersHandler, nil)
	users.RegisterHandlers(e, strictUsersHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
