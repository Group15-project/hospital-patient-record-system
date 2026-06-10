package models

import "gorm.io/gorm"

type RefreshToken struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null"`
	User   User

	Token string `gorm:"unique;not null"`

	ExpiresAt int64

	Revoked bool `gorm:"default:false"`

	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}