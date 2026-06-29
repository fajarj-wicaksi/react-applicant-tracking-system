package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ats-backend/internal/application/dto"
	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type CandidateHandler struct {
	candidateRepo   port.CandidateRepository
	applicationRepo port.ApplicationRepository
	positionRepo    port.PositionRepository
}

func NewCandidateHandler(cr port.CandidateRepository, ar port.ApplicationRepository, pr port.PositionRepository) *CandidateHandler {
	return &CandidateHandler{
		candidateRepo:   cr,
		applicationRepo: ar,
		positionRepo:    pr,
	}
}

// POST /api/v1/candidates
func (h *CandidateHandler) CreateCandidate(c *gin.Context) {
	var req dto.CreateCandidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	candidate := &domain.Candidate{
		TenantID:  tenantID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		ResumeURL: req.ResumeURL,
		Source:    req.Source,
	}

	if err := h.candidateRepo.Create(c.Request.Context(), candidate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create candidate"})
		return
	}

	c.JSON(http.StatusCreated, candidate)
}

// GET /api/v1/candidates
func (h *CandidateHandler) ListCandidates(c *gin.Context) {
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	candidates, err := h.candidateRepo.ListByTenant(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list candidates"})
		return
	}
	c.JSON(http.StatusOK, candidates)
}

// GET /api/v1/candidates/:id
func (h *CandidateHandler) GetCandidate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	candidate, err := h.candidateRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}
	if candidate.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// PUT /api/v1/candidates/:id
func (h *CandidateHandler) UpdateCandidate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	var req dto.UpdateCandidateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	candidate, err := h.candidateRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}
	if candidate.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if req.FirstName != "" {
		candidate.FirstName = req.FirstName
	}
	if req.LastName != "" {
		candidate.LastName = req.LastName
	}
	if req.Email != "" {
		candidate.Email = req.Email
	}
	if req.Phone != "" {
		candidate.Phone = req.Phone
	}
	if req.ResumeURL != "" {
		candidate.ResumeURL = req.ResumeURL
	}
	if req.Source != "" {
		candidate.Source = req.Source
	}

	if err := h.candidateRepo.Update(c.Request.Context(), candidate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update candidate"})
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// DELETE /api/v1/candidates/:id
func (h *CandidateHandler) DeleteCandidate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid candidate ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	candidate, err := h.candidateRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Candidate not found"})
		return
	}
	if candidate.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := h.candidateRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete candidate"})
		return
	}

	c.Status(http.StatusNoContent)
}

// POST /api/v1/positions
func (h *CandidateHandler) CreatePosition(c *gin.Context) {
	var req dto.CreatePositionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	position := &domain.Position{
		TenantID:    tenantID,
		Title:       req.Title,
		Department:  req.Department,
		Location:    req.Location,
		Type:        req.Type,
		Description: req.Description,
	}

	if err := h.positionRepo.Create(c.Request.Context(), position); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create position"})
		return
	}

	c.JSON(http.StatusCreated, position)
}

// GET /api/v1/positions
func (h *CandidateHandler) ListPositions(c *gin.Context) {
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	positions, err := h.positionRepo.ListByTenant(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list positions"})
		return
	}
	c.JSON(http.StatusOK, positions)
}

// GET /api/v1/positions/:id
func (h *CandidateHandler) GetPosition(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	position, err := h.positionRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return
	}
	if position.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	c.JSON(http.StatusOK, position)
}

// PUT /api/v1/positions/:id
func (h *CandidateHandler) UpdatePosition(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	var req dto.UpdatePositionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	position, err := h.positionRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return
	}
	if position.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if req.Title != "" {
		position.Title = req.Title
	}
	if req.Department != "" {
		position.Department = req.Department
	}
	if req.Location != "" {
		position.Location = req.Location
	}
	if req.Type != "" {
		position.Type = req.Type
	}
	if req.Description != "" {
		position.Description = req.Description
	}
	if req.IsOpen != nil {
		position.IsOpen = *req.IsOpen
	}

	if err := h.positionRepo.Update(c.Request.Context(), position); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update position"})
		return
	}

	c.JSON(http.StatusOK, position)
}

// DELETE /api/v1/positions/:id
func (h *CandidateHandler) DeletePosition(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position ID"})
		return
	}
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	position, err := h.positionRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return
	}
	if position.TenantID != tenantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := h.positionRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete position"})
		return
	}

	c.Status(http.StatusNoContent)
}

// POST /api/v1/applications
func (h *CandidateHandler) CreateApplication(c *gin.Context) {
	var req dto.CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenantID, _ := uuid.Parse(c.GetString("tenantID"))
	candidateID, _ := uuid.Parse(req.CandidateID)
	positionID, _ := uuid.Parse(req.PositionID)

	app := &domain.Application{
		TenantID:    tenantID,
		CandidateID: candidateID,
		PositionID:  positionID,
		Stage:       domain.StageApplied,
		Notes:       req.Notes,
	}

	if err := h.applicationRepo.Create(c.Request.Context(), app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
		return
	}

	c.JSON(http.StatusCreated, app)
}

// GET /api/v1/pipeline
func (h *CandidateHandler) GetPipeline(c *gin.Context) {
	tenantID, _ := uuid.Parse(c.GetString("tenantID"))

	grouped, err := h.applicationRepo.ListByTenantGrouped(c.Request.Context(), tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pipeline"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stages": grouped})
}

// PATCH /api/v1/applications/:id/stage
func (h *CandidateHandler) UpdateApplicationStage(c *gin.Context) {
	appID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	var req dto.UpdateStageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.applicationRepo.UpdateStage(c.Request.Context(), appID, domain.ApplicationStage(req.Stage), req.StageOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stage updated successfully"})
}
