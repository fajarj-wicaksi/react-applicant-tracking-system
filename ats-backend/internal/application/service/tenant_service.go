package service

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"ats-backend/internal/application/dto"
	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"
)

type tenantService struct {
	repo port.TenantRepository
}

// NewTenantService creates a new instance of TenantService
func NewTenantService(repo port.TenantRepository) port.TenantService {
	return &tenantService{repo: repo}
}

func mapToResponse(t *domain.Tenant) *dto.TenantResponse {
	return &dto.TenantResponse{
		ID:        t.ID,
		Name:      t.Name,
		Domain:    t.Domain,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func (s *tenantService) CreateTenant(ctx context.Context, req *dto.CreateTenantRequest) (*dto.TenantResponse, error) {
	existing, _ := s.repo.GetByDomain(ctx, req.Domain)
	if existing != nil {
		return nil, errors.New("domain already exists")
	}

	tenant := &domain.Tenant{
		Name:   req.Name,
		Domain: req.Domain,
	}

	if err := s.repo.Create(ctx, tenant); err != nil {
		return nil, err
	}

	return mapToResponse(tenant), nil
}

func (s *tenantService) GetTenantByID(ctx context.Context, id uuid.UUID) (*dto.TenantResponse, error) {
	tenant, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapToResponse(tenant), nil
}

func (s *tenantService) GetTenantByDomain(ctx context.Context, domain string) (*dto.TenantResponse, error) {
	tenant, err := s.repo.GetByDomain(ctx, domain)
	if err != nil {
		return nil, err
	}
	return mapToResponse(tenant), nil
}

func (s *tenantService) ListTenants(ctx context.Context, limit, offset int) ([]*dto.TenantResponse, error) {
	tenants, err := s.repo.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.TenantResponse, len(tenants))
	for i, t := range tenants {
		responses[i] = mapToResponse(t)
	}

	return responses, nil
}

func (s *tenantService) UpdateTenant(ctx context.Context, id uuid.UUID, req *dto.UpdateTenantRequest) (*dto.TenantResponse, error) {
	tenant, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		tenant.Name = req.Name
	}
	if req.Domain != "" {
		// Check uniqueness if domain changes
		if req.Domain != tenant.Domain {
			existing, _ := s.repo.GetByDomain(ctx, req.Domain)
			if existing != nil {
				return nil, errors.New("domain already exists")
			}
			tenant.Domain = req.Domain
		}
	}

	if err := s.repo.Update(ctx, tenant); err != nil {
		return nil, err
	}

	return mapToResponse(tenant), nil
}

func (s *tenantService) DeleteTenant(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
