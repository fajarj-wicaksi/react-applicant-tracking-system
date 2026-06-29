package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Position represents a job opening within a tenant
type Position struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	TenantID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"tenantId"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	Department  string         `gorm:"type:varchar(100)" json:"department"`
	Location    string         `gorm:"type:varchar(255)" json:"location"`
	Type        string         `gorm:"type:varchar(50)" json:"type"` // Full-time, Part-time, Contract
	Description string         `gorm:"type:text" json:"description"`
	IsOpen      bool           `gorm:"default:true" json:"isOpen"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Tenant *Tenant `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
}

func (p *Position) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
