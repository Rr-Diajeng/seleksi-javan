package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	Username  string    `gorm:"type:varchar;not_null;unique"`
	Email     string    `gorm:"type:varchar;not_null;unique"`
	Password  string    `gorm:"type:varchar;not_null"`
	Tasks     []Task    `gorm:"foreignKey:AssignedID"`
	CreatedAt time.Time `gorm:"type:timestamp;column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;column:updated_at"`
	DeletedAt gorm.DeletedAt
}
