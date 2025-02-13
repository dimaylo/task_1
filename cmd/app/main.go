package main

import (
	"log"

	"REST_API/internal/database"
	"REST_API/internal/handlers"
	"REST_API/internal/taskService"
	"REST_API/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)

	tasks.RegisterHandlers(e, strictHandler)

	log.Fatal(e.Start(":8080"))
}
