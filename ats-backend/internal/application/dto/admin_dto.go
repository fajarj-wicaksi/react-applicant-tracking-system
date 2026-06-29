package dto

// --- Admin Panel DTOs ---

type TenantAdminResponse struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	Domain         string                `json:"domain"`
	UserCount      int64                 `json:"userCount"`
	PositionCount  int64                 `json:"positionCount"`
	Subscription   *SubscriptionResponse `json:"subscription,omitempty"`
	CreatedAt      string                `json:"createdAt"`
}

type SubscriptionResponse struct {
	ID         string              `json:"id"`
	Status     string              `json:"status"`
	StartDate  string              `json:"startDate"`
	EndDate    string              `json:"endDate,omitempty"`
	Plan       *BillingPlanResponse `json:"plan,omitempty"`
}

type BillingPlanResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	PriceMonthly float64 `json:"priceMonthly"`
	MaxUsers     int     `json:"maxUsers"`
	MaxStorage   int64   `json:"maxStorage"`
	MaxPositions int     `json:"maxPositions"`
}

type SystemMonitoringResponse struct {
	TotalTenants       int64 `json:"totalTenants"`
	TotalUsers         int64 `json:"totalUsers"`
	TotalPositions     int64 `json:"totalPositions"`
	TotalApplications  int64 `json:"totalApplications"`
	ActiveSubscriptions int64 `json:"activeSubscriptions"`
}

type TenantStatsResponse struct {
	TenantID      string `json:"tenantId"`
	TenantName    string `json:"tenantName"`
	UserCount     int64  `json:"userCount"`
	PositionCount int64  `json:"positionCount"`
	ApplicationCount int64 `json:"applicationCount"`
	StorageUsage  int64  `json:"storageUsage"` // estimated bytes
}
