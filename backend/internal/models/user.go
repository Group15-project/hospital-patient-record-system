package models

import "gorm.io/gorm"

type User struct {
	ID uint `gorm:"primaryKey"`

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`

	Email string `gorm:"unique;not null"`

	Phone string

	HashedPassword string `gorm:"not null"`

	IsActive bool `gorm:"default:true"`

	RoleID uint
	Role   Role

	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}