package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type taskRepository struct{ db *gorm.DB }

func NewTaskRepository(db *gorm.DB) port.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(ctx context.Context, task *domain.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	var task domain.Task
	err := r.db.WithContext(ctx).First(&task, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) ListByTenant(ctx context.Context, tenantID uuid.UUID) ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.WithContext(ctx).
		Where("tenant_id = ?", tenantID).
		Order("created_at DESC").
		Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Update(ctx context.Context, task *domain.Task) error {
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *taskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Task{}, "id = ?", id).Error
}
