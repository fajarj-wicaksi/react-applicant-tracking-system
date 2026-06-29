package port

import (
	"context"

	"github.com/google/uuid"

	"ats-backend/internal/domain"
)

type InterviewRepository interface {
	Create(ctx context.Context, interview *domain.Interview) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Interview, error)
	ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]domain.Interview, error)
	Update(ctx context.Context, interview *domain.Interview) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
	ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}
