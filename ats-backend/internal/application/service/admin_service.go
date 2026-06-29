package service

import (
	"context"
	"fmt"

	"ats-backend/internal/application/dto"
	"ats-backend/internal/application/port"
	"ats-backend/internal/domain"

	"gorm.io/gorm"
)

type adminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) port.AdminService {
	return &adminService{db: db}
}

func (s *adminService) ListTenants(ctx context.Context) ([]dto.TenantAdminResponse, error) {
	var tenants []domain.Tenant
	if err := s.db.WithContext(ctx).Find(&tenants).Error; err != nil {
		return nil, err
	}

	var responses []dto.TenantAdminResponse
	for _, t := range tenants {
		var userCount int64
		s.db.WithContext(ctx).Model(&domain.User{}).Where("tenant_id = ?", t.ID).Count(&userCount)

		var posCount int64
		s.db.WithContext(ctx).Model(&domain.Position{}).Where("tenant_id = ?", t.ID).Count(&posCount)

		resp := dto.TenantAdminResponse{
			ID:            t.ID.String(),
			Name:          t.Name,
			Domain:        t.Domain,
			UserCount:     userCount,
			PositionCount: posCount,
			CreatedAt:     t.CreatedAt.Format("2006-01-02T15:04:05Z"),
		}

		// Try to find subscription
		var sub domain.Subscription
		if err := s.db.WithContext(ctx).Preload("BillingPlan").Where("tenant_id = ?", t.ID).First(&sub).Error; err == nil {
			subResp := &dto.SubscriptionResponse{
				ID:        sub.ID.String(),
				Status:    sub.Status,
				StartDate: sub.StartDate.Format("2006-01-02T15:04:05Z"),
			}
			if sub.EndDate != nil {
				end := sub.EndDate.Format("2006-01-02T15:04:05Z")
				subResp.EndDate = end
			}
			if sub.BillingPlan != nil {
				subResp.Plan = &dto.BillingPlanResponse{
					ID:           sub.BillingPlan.ID.String(),
					Name:         sub.BillingPlan.Name,
					Description:  sub.BillingPlan.Description,
					PriceMonthly: sub.BillingPlan.PriceMonthly,
					MaxUsers:     sub.BillingPlan.MaxUsers,
					MaxStorage:   sub.BillingPlan.MaxStorage,
					MaxPositions: sub.BillingPlan.MaxPositions,
				}
			}
			resp.Subscription = subResp
		}

		responses = append(responses, resp)
	}
	return responses, nil
}

func (s *adminService) GetTenantStats(ctx context.Context, tenantID string) (*dto.TenantStatsResponse, error) {
	var tenant domain.Tenant
	if err := s.db.WithContext(ctx).First(&tenant, "id = ?", tenantID).Error; err != nil {
		return nil, fmt.Errorf("tenant not found")
	}

	var userCount, posCount, appCount int64
	s.db.WithContext(ctx).Model(&domain.User{}).Where("tenant_id = ?", tenantID).Count(&userCount)
	s.db.WithContext(ctx).Model(&domain.Position{}).Where("tenant_id = ?", tenantID).Count(&posCount)
	s.db.WithContext(ctx).Model(&domain.Application{}).Where("tenant_id = ?", tenantID).Count(&appCount)

	// Estimate storage: ~2KB per record as approximation
	storageUsage := (userCount + posCount + appCount) * 2048

	return &dto.TenantStatsResponse{
		TenantID:         tenantID,
		TenantName:       tenant.Name,
		UserCount:        userCount,
		PositionCount:    posCount,
		ApplicationCount: appCount,
		StorageUsage:     storageUsage,
	}, nil
}

func (s *adminService) GetSystemMonitoring(ctx context.Context) (*dto.SystemMonitoringResponse, error) {
	var totalTenants, totalUsers, totalPositions, totalApps, activeSubs int64
	s.db.WithContext(ctx).Model(&domain.Tenant{}).Count(&totalTenants)
	s.db.WithContext(ctx).Model(&domain.User{}).Count(&totalUsers)
	s.db.WithContext(ctx).Model(&domain.Position{}).Count(&totalPositions)
	s.db.WithContext(ctx).Model(&domain.Application{}).Count(&totalApps)
	s.db.WithContext(ctx).Model(&domain.Subscription{}).Where("status = ?", "active").Count(&activeSubs)

	return &dto.SystemMonitoringResponse{
		TotalTenants:        totalTenants,
		TotalUsers:          totalUsers,
		TotalPositions:      totalPositions,
		TotalApplications:   totalApps,
		ActiveSubscriptions: activeSubs,
	}, nil
}
