package model

import (
	"time"

	"gorm.io/gorm"
)

type StatusTask string

const (
	Pending    StatusTask = "pending"
	InProgress StatusTask = "in_progress"
	Completed  StatusTask = "completed"
)

type Task struct {
	ID          uint       `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	Title       string     `gorm:"type:varchar;not_null"`
	Description string     `gorm:"type:text"`
	Status      StatusTask `gorm:"type:status;default:'pending';not_null"`
	AssignedID  uint       `gorm:"type:bigint;not_null"`
	CreatedAt   time.Time  `gorm:"type:timestamp;column:created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;column:updated_at"`
	DeletedAt   gorm.DeletedAt
}
