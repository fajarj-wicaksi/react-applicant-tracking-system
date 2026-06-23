package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateTenantRequest represents the payload to create a new tenant
type CreateTenantRequest struct {
	Name   string `json:"name" binding:"required,min=2,max=255"`
	Domain string `json:"domain" binding:"required,fqdn|hostname"`
}

// UpdateTenantRequest represents the payload to update a tenant
type UpdateTenantRequest struct {
	Name   string `json:"name" binding:"omitempty,min=2,max=255"`
	Domain string `json:"domain" binding:"omitempty,fqdn|hostname"`
}

// TenantResponse represents the tenant data returned to the client
type TenantResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
