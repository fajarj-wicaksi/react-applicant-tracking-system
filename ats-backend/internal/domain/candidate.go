package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ApplicationStage represents the current stage of a candidate's application
type ApplicationStage string

const (
	StageApplied   ApplicationStage = "Applied"
	StageScreening ApplicationStage = "Screening"
	StageInterview ApplicationStage = "Interview"
	StageOffer     ApplicationStage = "Offer"
	StageHired     ApplicationStage = "Hired"
	StageRejected  ApplicationStage = "Rejected"
)

// Candidate represents a person applying for positions
type Candidate struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	TenantID  uuid.UUID      `gorm:"type:uuid;not null;index" json:"tenantId"`
	FirstName string         `gorm:"type:varchar(100);not null" json:"firstName"`
	LastName  string         `gorm:"type:varchar(100);not null" json:"lastName"`
	Email     string         `gorm:"type:varchar(255);not null;uniqueIndex:idx_candidate_email_tenant" json:"email"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	ResumeURL string         `gorm:"type:varchar(500)" json:"resumeUrl"`
	Source    string         `gorm:"type:varchar(100)" json:"source"` // LinkedIn, Referral, Website, etc.
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`

	Tenant       *Tenant       `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	Applications []Application `gorm:"foreignKey:CandidateID" json:"applications,omitempty"`
}

func (c *Candidate) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

// Application represents a candidate's application to a specific position
type Application struct {
	ID          uuid.UUID        `gorm:"type:uuid;primaryKey" json:"id"`
	TenantID    uuid.UUID        `gorm:"type:uuid;not null;index" json:"tenantId"`
	CandidateID uuid.UUID        `gorm:"type:uuid;not null;index" json:"candidateId"`
	PositionID  uuid.UUID        `gorm:"type:uuid;not null;index" json:"positionId"`
	Stage       ApplicationStage `gorm:"type:varchar(50);not null;default:'Applied'" json:"stage"`
	StageOrder  int              `gorm:"default:0" json:"stageOrder"` // ordering within the same stage column
	Notes       string           `gorm:"type:text" json:"notes"`
	AppliedAt   time.Time        `gorm:"autoCreateTime" json:"appliedAt"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"deletedAt,omitempty"`

	Tenant    *Tenant    `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
	Candidate *Candidate `gorm:"foreignKey:CandidateID" json:"candidate,omitempty"`
	Position  *Position  `gorm:"foreignKey:PositionID" json:"position,omitempty"`
}

func (a *Application) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
