package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"task-api/models"
	"task-api/store"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store *store.Store
}

func New(store *store.Store) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) GetTasksHandler(c *gin.Context) {
	tasks, err := h.store.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) CreateTasksHandler(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title is required",
		})
		return
	}

	if task.Status == "" {
		task.Status = "pending"
	}

	if task.Status != "pending" &&
		task.Status != "in_progress" &&
		task.Status != "done" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid status",
		})
		return
	}

	task.CreatedAt = time.Now()

	err = h.store.Create(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create task",
		})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *Handler) GetTasksByIdHandler(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	task, err := h.store.GetByID(idNum)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTasksByIdHandler(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	err = h.store.DeleteByID(idNum)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}

func (h *Handler) UpdateTasksByIdHandler(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var task models.Task

	err = c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	if task.Status != "" {
		if task.Status != "pending" &&
			task.Status != "in_progress" &&
			task.Status != "done" {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid status",
			})
			return
		}
	}

	err = h.store.UpdateByID(idNum, task)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	updatedTask, err := h.store.GetByID(idNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Task updated successfully",
		})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}
