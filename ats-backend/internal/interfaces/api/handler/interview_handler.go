package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type InterviewHandler struct {
	repo port.InterviewRepository
}

func NewInterviewHandler(repo port.InterviewRepository) *InterviewHandler {
	return &InterviewHandler{repo: repo}
}

type createInterviewReq struct {
	CandidateID string `json:"candidateId" binding:"required"`
	PositionID  string `json:"positionId" binding:"required"`
	Title       string `json:"title" binding:"required"`
	ScheduledAt string `json:"scheduledAt" binding:"required"`
	Duration    int    `json:"duration"`
	Location    string `json:"location"`
	Notes       string `json:"notes"`
}

type updateInterviewReq struct {
	Title       string `json:"title"`
	ScheduledAt string `json:"scheduledAt"`
	Duration    int    `json:"duration"`
	Location    string `json:"location"`
	Notes       string `json:"notes"`
	Status      string `json:"status"`
}

// POST /api/v1/interviews
func (h *InterviewHandler) CreateInterview(c *gin.Context) {
	var req createInterviewReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := uuid.Parse(c.GetString("tenantID"))
	candidateID, _ := uuid.Parse(req.CandidateID)
	positionID, _ := uuid.Parse(req.PositionID)
	scheduledAt, err := time.Parse(time.RFC3339, req.ScheduledAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scheduledAt format. Use ISO8601 (RFC3339)"})
		return
	}

	duration := req.Duration
	if duration <= 0 {
		duration = 60
	}

	interview := &domain.Interview{
		TenantID:    tenantID,
		CandidateID: candidateID,
		PositionID:  positionID,
		Title:       req.Title,
		ScheduledAt: scheduledAt,
		Duration:    duration,
		Location:    req.Location,
		Notes:       req.Notes,
		Status:      domain.InterviewScheduled,
	}

	if err := h.repo.Create(c.Request.Context(), interview); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create interview"})
		return
	}

	c.JSON(http.StatusCreated, interview)
}

// GET /api/v1/interviews
func (h *InterviewHandler) ListInterviews(c *gin.Context) {
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))
	interviews, err := h.repo.ListByTenant(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list interviews"})
		return
	}
	c.JSON(http.StatusOK, interviews)
}

// GET /api/v1/interviews/:id
func (h *InterviewHandler) GetInterview(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interview ID"})
		return
	}
	interview, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Interview not found"})
		return
	}
	c.JSON(http.StatusOK, interview)
}

// PUT /api/v1/interviews/:id
func (h *InterviewHandler) UpdateInterview(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interview ID"})
		return
	}

	interview, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Interview not found"})
		return
	}

	var req updateInterviewReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Title != "" {
		interview.Title = req.Title
	}
	if req.ScheduledAt != "" {
		t, err := time.Parse(time.RFC3339, req.ScheduledAt)
		if err == nil {
			interview.ScheduledAt = t
		}
	}
	if req.Duration > 0 {
		interview.Duration = req.Duration
	}
	if req.Location != "" {
		interview.Location = req.Location
	}
	if req.Notes != "" {
		interview.Notes = req.Notes
	}
	if req.Status != "" {
		interview.Status = domain.InterviewStatus(req.Status)
	}

	if err := h.repo.Update(c.Request.Context(), interview); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update interview"})
		return
	}

	c.JSON(http.StatusOK, interview)
}

// DELETE /api/v1/interviews/:id
func (h *InterviewHandler) DeleteInterview(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interview ID"})
		return
	}
	if err := h.repo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete interview"})
		return
	}
	c.Status(http.StatusNoContent)
}
