package port

import (
	"context"

	"github.com/google/uuid"

	"ats-backend/internal/domain"
)

type RoleRepository interface {
	Create(ctx context.Context, role *domain.Role) error
	GetByName(ctx context.Context, name string, tenantID *uuid.UUID) (*domain.Role, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Role, error)
}
