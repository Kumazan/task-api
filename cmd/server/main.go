package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/kumazan/task-api/internal/api"
	"github.com/kumazan/task-api/internal/repository"
	"github.com/kumazan/task-api/internal/service"
)

func main() {
	r := gin.Default()

	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := api.NewTaskHandler(taskService)

	r.GET("/tasks", taskHandler.ListTasks)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	log.Println("Server starting on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
