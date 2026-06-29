package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type interviewRepository struct{ db *gorm.DB }

func NewInterviewRepository(db *gorm.DB) port.InterviewRepository {
	return &interviewRepository{db: db}
}

func (r *interviewRepository) Create(ctx context.Context, interview *domain.Interview) error {
	return r.db.WithContext(ctx).Create(interview).Error
}

func (r *interviewRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Interview, error) {
	var interview domain.Interview
	err := r.db.WithContext(ctx).Preload("Candidate").Preload("Position").First(&interview, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &interview, nil
}

func (r *interviewRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]domain.Interview, error) {
	var interviews []domain.Interview
	err := r.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Preload("Candidate").
		Preload("Position").
		Order("scheduled_at ASC").
		Find(&interviews).Error
	return interviews, err
}

func (r *interviewRepository) Update(ctx context.Context, interview *domain.Interview) error {
	return r.db.WithContext(ctx).Save(interview).Error
}

func (r *interviewRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Interview{}, "id = ?", id).Error
}
