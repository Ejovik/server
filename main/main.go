package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func GETHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func POSTHandler(w http.ResponseWriter, r *http.Request) {
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

func PATCHHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var reqBody requestBody
	json.NewDecoder(r.Body).Decode(&reqBody)

	var task Task

	if err := DB.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.Task = reqBody.Task
	task.IsDone = reqBody.IsDone

	DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func DELETEHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task Task
	if err := DB.First(&task, id).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	DB.Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GETHandler).Methods("GET")
	router.HandleFunc("/api/tasks", POSTHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", PATCHHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DELETEHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
