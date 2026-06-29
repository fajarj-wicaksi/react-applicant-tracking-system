package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type applicationRepository struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) port.ApplicationRepository {
	return &applicationRepository{db: db}
}

func (r *applicationRepository) Create(ctx context.Context, app *domain.Application) error {
	return r.db.WithContext(ctx).Create(app).Error
}

func (r *applicationRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Application, error) {
	var app domain.Application
	err := r.db.WithContext(ctx).Preload("Candidate").Preload("Position").First(&app, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *applicationRepository) UpdateStage(ctx context.Context, id uuid.UUID, stage domain.ApplicationStage, order int) error {
	return r.db.WithContext(ctx).Model(&domain.Application{}).Where("id = ?", id).Updates(map[string]interface{}{
		"stage":       stage,
		"stage_order": order,
	}).Error
}

func (r *applicationRepository) ListByTenantGrouped(ctx context.Context, tenantID uuid.UUID) (map[domain.ApplicationStage][]domain.Application, error) {
	var apps []domain.Application
	err := r.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Preload("Candidate").
		Preload("Position").
		Order("stage_order ASC").
		Find(&apps).Error
	if err != nil {
		return nil, err
	}

	// Group by stage
	grouped := make(map[domain.ApplicationStage][]domain.Application)
	stages := []domain.ApplicationStage{
		domain.StageApplied, domain.StageScreening, domain.StageInterview,
		domain.StageOffer, domain.StageHired, domain.StageRejected,
	}
	for _, s := range stages {
		grouped[s] = []domain.Application{}
	}
	for _, app := range apps {
		grouped[app.Stage] = append(grouped[app.Stage], app)
	}
	return grouped, nil
}
