package main

import (
	"net/http"
	"task-api/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", handlers.GetTasksHandler)
	mux.HandleFunc("POST /tasks", handlers.CreateTasksHandler)
	mux.HandleFunc("GET /tasks/{id}", handlers.GetTasksByIdHandler)
	mux.HandleFunc("DELETE /tasks/{id}", handlers.DeleteTasksByIdHandler)
	mux.HandleFunc("PUT /tasks/{id}", handlers.UpdateTasksByIdHandler)

	http.ListenAndServe(":8080", mux)
}
