package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskPending    TaskStatus = "Pending"
	TaskInProgress TaskStatus = "In Progress"
	TaskCompleted  TaskStatus = "Completed"
)

type Task struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key"`
	TenantID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	DueDate     *time.Time `gorm:"index"`
	Status      TaskStatus `gorm:"type:varchar(50);not null;default:'Pending'"`
	AssignedTo  *uuid.UUID `gorm:"type:uuid;index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Tenant   *Tenant `gorm:"foreignKey:TenantID"`
	Assignee *User   `gorm:"foreignKey:AssignedTo"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}
