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

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Task: %s", task)
	} else {
		fmt.Fprintln(w, "not method:")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	task = reqBody.Message
	fmt.Fprintln(w, "Task updated:", task)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Привет, Мир!")
	} else {
		fmt.Fprintln(w, "Неправильный метод")
	}

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/task", GetHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)

}
