package domain

import (
	"time"

	"github.com/google/uuid"
)

// Tenant represents an organization or company using the ATS
type Tenant struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Domain    string    `gorm:"type:varchar(255);uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
