package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	TenantID    *uuid.UUID     `gorm:"type:uuid;index" json:"tenantId,omitempty"` // nil for system roles
	Name        string         `gorm:"type:varchar(50);not null" json:"name"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	Permissions []string       `gorm:"type:text[];serializer:json" json:"permissions"` // simple array of permission strings
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Tenant *Tenant `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
