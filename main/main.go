package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	json.NewDecoder(r.Body).Decode(&reqBody)
	task = reqBody.Message
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", UpdateTaskHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
