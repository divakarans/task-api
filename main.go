package main

import (
	"log"
	"task-api/db"
	"task-api/handlers"
	"task-api/store"

	"github.com/gin-gonic/gin"
)

func main() {

	database, err := db.InitDB("tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	taskStore := store.New(database)

	handler := handlers.New(taskStore)

	r := gin.Default()

	r.GET("/tasks", handler.GetTasksHandler)
	r.POST("/tasks", handler.CreateTasksHandler)
	r.GET("/tasks/:id", handler.GetTasksByIdHandler)
	r.PUT("/tasks/:id", handler.UpdateTasksByIdHandler)
	r.DELETE("/tasks/:id", handler.DeleteTasksByIdHandler)

	r.Run(":8080")
}
