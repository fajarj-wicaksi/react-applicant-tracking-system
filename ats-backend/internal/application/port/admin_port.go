package port

import (
	"context"

	"ats-backend/internal/application/dto"
)

type AdminService interface {
	ListTenants(ctx context.Context) ([]dto.TenantAdminResponse, error)
	GetTenantStats(ctx context.Context, tenantID string) (*dto.TenantStatsResponse, error)
	GetSystemMonitoring(ctx context.Context) (*dto.SystemMonitoringResponse, error)
}
