package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type candidateRepository struct {
	db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) port.CandidateRepository {
	return &candidateRepository{db: db}
}

func (r *candidateRepository) Create(ctx context.Context, candidate *domain.Candidate) error {
	return r.db.WithContext(ctx).Create(candidate).Error
}

func (r *candidateRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Candidate, error) {
	var candidate domain.Candidate
	err := r.db.WithContext(ctx).Preload("Applications.Position").First(&candidate, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (r *candidateRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]domain.Candidate, error) {
	var candidates []domain.Candidate
	err := r.db.WithContext(ctx).Where("tenant_id = ?", tenantID).Order("created_at DESC").Find(&candidates).Error
	return candidates, err
}

func (r *candidateRepository) Update(ctx context.Context, candidate *domain.Candidate) error {
	return r.db.WithContext(ctx).Save(candidate).Error
}

func (r *candidateRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Candidate{}, "id = ?", id).Error
}
