package dto

import "ats-backend/internal/domain"

// --- Candidate DTOs ---

type CreateCandidateRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	ResumeURL string `json:"resumeUrl"`
	Source    string `json:"source"`
}

type UpdateCandidateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ResumeURL string `json:"resumeUrl"`
	Source    string `json:"source"`
}

type CandidateResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	ResumeURL string `json:"resumeUrl"`
	Source    string `json:"source"`
	CreatedAt string `json:"createdAt"`
}

// --- Application / Pipeline DTOs ---

type CreateApplicationRequest struct {
	CandidateID string `json:"candidateId" binding:"required"`
	PositionID  string `json:"positionId" binding:"required"`
	Notes       string `json:"notes"`
}

type UpdateStageRequest struct {
	Stage      string `json:"stage" binding:"required"`
	StageOrder int    `json:"stageOrder"`
}

type ApplicationResponse struct {
	ID          string             `json:"id"`
	CandidateID string             `json:"candidateId"`
	PositionID  string             `json:"positionId"`
	Stage       string             `json:"stage"`
	StageOrder  int                `json:"stageOrder"`
	Notes       string             `json:"notes"`
	AppliedAt   string             `json:"appliedAt"`
	Candidate   *CandidateResponse `json:"candidate,omitempty"`
	Position    *PositionResponse  `json:"position,omitempty"`
}

// --- Position DTOs ---

type CreatePositionRequest struct {
	Title       string `json:"title" binding:"required"`
	Department  string `json:"department"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type UpdatePositionRequest struct {
	Title       string `json:"title"`
	Department  string `json:"department"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Description string `json:"description"`
	IsOpen      *bool  `json:"isOpen"`
}

type PositionResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Department  string `json:"department"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Description string `json:"description"`
	IsOpen      bool   `json:"isOpen"`
	CreatedAt   string `json:"createdAt"`
}

// --- Pipeline (Kanban) DTO ---

type PipelineResponse struct {
	Stages map[domain.ApplicationStage][]ApplicationResponse `json:"stages"`
}
