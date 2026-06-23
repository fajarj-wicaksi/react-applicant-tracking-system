package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) port.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(ctx context.Context, role *domain.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *roleRepository) GetByName(ctx context.Context, name string, tenantID *uuid.UUID) (*domain.Role, error) {
	var role domain.Role
	query := r.db.WithContext(ctx).Where("name = ?", name)
	if tenantID != nil {
		query = query.Where("tenant_id = ? OR tenant_id IS NULL", tenantID)
	} else {
		query = query.Where("tenant_id IS NULL")
	}
	err := query.First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	var role domain.Role
	err := r.db.WithContext(ctx).First(&role, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
