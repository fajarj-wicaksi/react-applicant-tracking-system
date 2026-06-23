package port

import (
	"context"

	"github.com/google/uuid"
	"ats-backend/internal/application/dto"
	"ats-backend/internal/domain"
)

// TenantRepository defines the interface for data access
type TenantRepository interface {
	Create(ctx context.Context, tenant *domain.Tenant) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Tenant, error)
	GetByDomain(ctx context.Context, domain string) (*domain.Tenant, error)
	GetAll(ctx context.Context, limit, offset int) ([]*domain.Tenant, error)
	Update(ctx context.Context, tenant *domain.Tenant) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TenantService defines the interface for business logic
type TenantService interface {
	CreateTenant(ctx context.Context, req *dto.CreateTenantRequest) (*dto.TenantResponse, error)
	GetTenantByID(ctx context.Context, id uuid.UUID) (*dto.TenantResponse, error)
	GetTenantByDomain(ctx context.Context, domain string) (*dto.TenantResponse, error)
	ListTenants(ctx context.Context, limit, offset int) ([]*dto.TenantResponse, error)
	UpdateTenant(ctx context.Context, id uuid.UUID, req *dto.UpdateTenantRequest) (*dto.TenantResponse, error)
	DeleteTenant(ctx context.Context, id uuid.UUID) error
}
