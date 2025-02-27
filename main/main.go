package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	json.NewDecoder(r.Body).Decode(&reqBody)

	newTask := Task{
		Task:   reqBody.Task,
		IsDone: reqBody.IsDone,
	}

	DB.Create(&newTask)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", HelloHandler).Methods("GET")
	router.HandleFunc("/api/tasks", UpdateTaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
