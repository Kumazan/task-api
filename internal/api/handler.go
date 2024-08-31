package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kumazan/task-api/internal/model"
	"github.com/kumazan/task-api/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req struct {
		Name   string `json:"name" binding:"required"`
		Status int    `json:"status" binding:"min=0,max=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.CreateTask(req.Name, model.TaskStatus(req.Status))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Name   string `json:"name"`
		Status int    `json:"status" binding:"min=0,max=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateTask(id, req.Name, model.TaskStatus(req.Status)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
