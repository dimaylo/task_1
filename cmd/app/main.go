package main

import (
	"log"
	"net/http"

	"REST_API/internal/database"
	"REST_API/internal/handlers"
	"REST_API/internal/taskService"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
