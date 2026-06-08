package models

import (
	"time"

	"gorm.io/gorm"
)

type Medication struct {
	ID uint `gorm:"primaryKey"`

	Name string `gorm:"uniqueIndex;not null"`

	GenericName string

	Strength string

	Form string
	// Tablet, Syrup, Injection, Capsule

	UnitPrice float64

	QuantityInStock int

	ReorderLevel int

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}