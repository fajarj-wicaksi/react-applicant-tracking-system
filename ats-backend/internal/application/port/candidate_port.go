package port

import (
	"context"

	"github.com/google/uuid"

	"ats-backend/internal/domain"
)

type CandidateRepository interface {
	Create(ctx context.Context, candidate *domain.Candidate) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Candidate, error)
	ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]domain.Candidate, error)
	Update(ctx context.Context, candidate *domain.Candidate) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ApplicationRepository interface {
	Create(ctx context.Context, app *domain.Application) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Application, error)
	UpdateStage(ctx context.Context, id uuid.UUID, stage domain.ApplicationStage, order int) error
	ListByTenantGrouped(ctx context.Context, tenantID uuid.UUID) (map[domain.ApplicationStage][]domain.Application, error)
}

type PositionRepository interface {
	Create(ctx context.Context, position *domain.Position) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Position, error)
	ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]domain.Position, error)
	Update(ctx context.Context, position *domain.Position) error
	Delete(ctx context.Context, id uuid.UUID) error
}
