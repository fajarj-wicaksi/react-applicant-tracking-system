package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InterviewStatus string

const (
	InterviewScheduled InterviewStatus = "Scheduled"
	InterviewCompleted InterviewStatus = "Completed"
	InterviewCancelled InterviewStatus = "Cancelled"
)

type Interview struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key"`
	TenantID    uuid.UUID       `gorm:"type:uuid;not null;index"`
	CandidateID uuid.UUID       `gorm:"type:uuid;not null;index"`
	PositionID  uuid.UUID       `gorm:"type:uuid;not null;index"`
	Title       string          `gorm:"type:varchar(255);not null"`
	ScheduledAt time.Time       `gorm:"not null"`
	Duration    int             `gorm:"not null;default:60"` // in minutes
	Status      InterviewStatus `gorm:"type:varchar(50);not null;default:'Scheduled'"`
	Location    string          `gorm:"type:varchar(255)"` // link or physical location
	Notes       string          `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Tenant    *Tenant    `gorm:"foreignKey:TenantID"`
	Candidate *Candidate `gorm:"foreignKey:CandidateID"`
	Position  *Position  `gorm:"foreignKey:PositionID"`
}

func (i *Interview) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return
}
