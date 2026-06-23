package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type tenantRepository struct {
	db *gorm.DB
}

// NewTenantRepository creates a new instance of TenantRepository
func NewTenantRepository(db *gorm.DB) port.TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	return r.db.WithContext(ctx).Create(tenant).Error
}

func (r *tenantRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Tenant, error) {
	var tenant domain.Tenant
	err := r.db.WithContext(ctx).First(&tenant, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) GetByDomain(ctx context.Context, domainName string) (*domain.Tenant, error) {
	var tenant domain.Tenant
	err := r.db.WithContext(ctx).First(&tenant, "domain = ?", domainName).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) GetAll(ctx context.Context, limit, offset int) ([]*domain.Tenant, error) {
	var tenants []*domain.Tenant
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Order("created_at DESC").Find(&tenants).Error
	return tenants, err
}

func (r *tenantRepository) Update(ctx context.Context, tenant *domain.Tenant) error {
	return r.db.WithContext(ctx).Save(tenant).Error
}

func (r *tenantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Tenant{}, "id = ?", id).Error
}
