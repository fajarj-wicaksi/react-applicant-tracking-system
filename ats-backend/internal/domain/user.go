package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	TenantID  uuid.UUID      `gorm:"type:uuid;not null;index" json:"tenantId"`
	RoleID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"roleId"`
	Email     string         `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"` // never serialize
	FirstName string         `gorm:"type:varchar(100);not null" json:"firstName"`
	LastName  string         `gorm:"type:varchar(100);not null" json:"lastName"`
	IsActive  bool           `gorm:"default:true" json:"isActive"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Tenant *Tenant `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	Role   *Role   `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
