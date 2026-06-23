package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ats-backend/internal/application/dto"
	"ats-backend/internal/domain"
)

type MockTenantRepository struct {
	mock.Mock
}

func (m *MockTenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	args := m.Called(ctx, tenant)
	return args.Error(0)
}

func (m *MockTenantRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Tenant, error) {
	args := m.Called(ctx, id)
	if t := args.Get(0); t != nil {
		return t.(*domain.Tenant), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenantRepository) GetByDomain(ctx context.Context, domainStr string) (*domain.Tenant, error) {
	args := m.Called(ctx, domainStr)
	if t := args.Get(0); t != nil {
		return t.(*domain.Tenant), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockTenantRepository) GetAll(ctx context.Context, limit, offset int) ([]*domain.Tenant, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*domain.Tenant), args.Error(1)
}

func (m *MockTenantRepository) Update(ctx context.Context, tenant *domain.Tenant) error {
	args := m.Called(ctx, tenant)
	return args.Error(0)
}

func (m *MockTenantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateTenant_Success(t *testing.T) {
	mockRepo := new(MockTenantRepository)
	svc := NewTenantService(mockRepo)

	req := &dto.CreateTenantRequest{
		Name:   "Test Corp",
		Domain: "testcorp.com",
	}

	mockRepo.On("GetByDomain", mock.Anything, "testcorp.com").Return(nil, errors.New("not found"))
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Tenant")).Return(nil)

	resp, err := svc.CreateTenant(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Test Corp", resp.Name)
	assert.Equal(t, "testcorp.com", resp.Domain)
	mockRepo.AssertExpectations(t)
}

func TestCreateTenant_DomainExists(t *testing.T) {
	mockRepo := new(MockTenantRepository)
	svc := NewTenantService(mockRepo)

	req := &dto.CreateTenantRequest{
		Name:   "Test Corp",
		Domain: "testcorp.com",
	}

	existing := &domain.Tenant{
		ID:     uuid.New(),
		Name:   "Existing",
		Domain: "testcorp.com",
	}

	mockRepo.On("GetByDomain", mock.Anything, "testcorp.com").Return(existing, nil)

	resp, err := svc.CreateTenant(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "domain already exists", err.Error())
	mockRepo.AssertExpectations(t)
}
