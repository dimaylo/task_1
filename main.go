package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Task, %s!", task)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	task = reqBody.Message
	fmt.Fprintln(w, "Task, updated!", task)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	result := DB.Create(&msg)
	if result.Error != nil {
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	result := DB.Find(&messages)
	if result.Error != nil {
		http.Error(w, "Failed go get message", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var msg Message
	result := DB.First(&msg, id)
	if result.Error != nil {
		http.Error(w, "Failed go get message", http.StatusInternalServerError)
		return
	}
	var updatedMessage Message
	err := json.NewDecoder(r.Body).Decode(&updatedMessage)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	msg.Task = updatedMessage.Task
	msg.IsDone = updatedMessage.IsDone
	DB.Save(&msg)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var msg Message
	result := DB.First(&msg, id)
	if result.Error != nil {
		http.Error(w, "Failed go get message", http.StatusInternalServerError)
		return
	}
	DB.Delete(&msg)
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", GetHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	router.HandleFunc("/api/messages", GetMessage).Methods("GET")
	router.HandleFunc("/api/messages/{id}", UpdateMessage).Methods("PUT")
	router.HandleFunc("/api/messages/{id}", DeleteMessage).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
