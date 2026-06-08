package models

import "gorm.io/gorm"

type AuditLog struct {
	ID uint `gorm:"primaryKey"`

	UserID *uint
	User   *User

	Action string `gorm:"not null"`

	Resource string `gorm:"not null"`

	ResourceID string

	IPAddress string

	UserAgent string

	CreatedAt int64

	DeletedAt gorm.DeletedAt `gorm:"index"`
}