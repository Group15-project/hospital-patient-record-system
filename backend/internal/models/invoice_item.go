package models

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceItem struct {
	ID uint `gorm:"primaryKey"`

	InvoiceID uint

	Description string

	Quantity int

	UnitPrice float64

	Amount float64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}