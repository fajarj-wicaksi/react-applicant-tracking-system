package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BillingPlan represents a pricing tier (Free, Pro, Enterprise)
type BillingPlan struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name          string         `gorm:"type:varchar(50);not null;uniqueIndex" json:"name"`
	Description   string         `gorm:"type:varchar(500)" json:"description"`
	PriceMonthly  float64        `gorm:"type:decimal(10,2);not null;default:0" json:"priceMonthly"`
	MaxUsers      int            `gorm:"not null;default:5" json:"maxUsers"`
	MaxStorage    int64          `gorm:"not null;default:1073741824" json:"maxStorage"` // bytes, default 1GB
	MaxPositions  int            `gorm:"not null;default:10" json:"maxPositions"`
	IsActive      bool           `gorm:"default:true" json:"isActive"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

func (b *BillingPlan) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}

// Subscription links a Tenant to a BillingPlan
type Subscription struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	TenantID      uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex" json:"tenantId"`
	BillingPlanID uuid.UUID      `gorm:"type:uuid;not null;index" json:"billingPlanId"`
	Status        string         `gorm:"type:varchar(20);not null;default:'active'" json:"status"` // active, expired, cancelled
	StartDate     time.Time      `gorm:"not null" json:"startDate"`
	EndDate       *time.Time     `json:"endDate,omitempty"` // nil = no expiry
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Tenant      *Tenant      `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	BillingPlan *BillingPlan `gorm:"foreignKey:BillingPlanID" json:"billingPlan,omitempty"`
}

func (s *Subscription) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}
