package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type positionRepository struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) port.PositionRepository {
	return &positionRepository{db: db}
}

func (r *positionRepository) Create(ctx context.Context, position *domain.Position) error {
	return r.db.WithContext(ctx).Create(position).Error
}

func (r *positionRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Position, error) {
	var position domain.Position
	err := r.db.WithContext(ctx).First(&position, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *positionRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]domain.Position, error) {
	var positions []domain.Position
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("created_at DESC").Find(&positions).Error
	return positions, err
}

func (r *positionRepository) Update(ctx context.Context, position *domain.Position) error {
	return r.db.WithContext(ctx).Save(position).Error
}

func (r *positionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Position{}, "id = ?", id).Error
}
