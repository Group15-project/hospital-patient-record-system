package models

import (
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	ID uint `gorm:"primaryKey"`

	UserID *uint
	User   *User `gorm:"foreignKey:UserID"`

	Action string `gorm:"size:100;not null"`

	Resource string `gorm:"size:100;not null"`

	ResourceID string `gorm:"size:100"`

	Details string `gorm:"type:text"`

	IPAddress string

	UserAgent string

	CreatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}