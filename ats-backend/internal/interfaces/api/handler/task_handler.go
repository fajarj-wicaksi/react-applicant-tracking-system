package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type TaskHandler struct {
	repo port.TaskRepository
}

func NewTaskHandler(repo port.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

type createTaskReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Status      string `json:"status"`
	AssignedTo  string `json:"assignedTo"`
}

type updateTaskReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Status      string `json:"status"`
	AssignedTo  string `json:"assignedTo"`
}

// POST /api/v1/tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req createTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	status := domain.TaskPending
	if req.Status != "" {
		status = domain.TaskStatus(req.Status)
	}

	task := &domain.Task{
		TenantID:    tenantID,
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
	}

	if req.DueDate != "" {
		t, err := time.Parse("2006-01-02", req.DueDate)
		if err == nil {
			task.DueDate = &t
		}
	}
	if req.AssignedTo != "" {
		uid, err := uuid.Parse(req.AssignedTo)
		if err == nil {
			task.AssignedTo = &uid
		}
	}

	if err := h.repo.Create(c.Request.Context(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GET /api/v1/tasks
func (h *TaskHandler) ListTasks(c *gin.Context) {
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))
	tasks, err := h.repo.ListByTenant(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GET /api/v1/tasks/:id
func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// PUT /api/v1/tasks/:id
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var req updateTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.Status != "" {
		task.Status = domain.TaskStatus(req.Status)
	}
	if req.DueDate != "" {
		t, err := time.Parse("2006-01-02", req.DueDate)
		if err == nil {
			task.DueDate = &t
		}
	}
	if req.AssignedTo != "" {
		uid, err := uuid.Parse(req.AssignedTo)
		if err == nil {
			task.AssignedTo = &uid
		}
	}

	if err := h.repo.Update(c.Request.Context(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// DELETE /api/v1/tasks/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if err := h.repo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	c.Status(http.StatusNoContent)
}
