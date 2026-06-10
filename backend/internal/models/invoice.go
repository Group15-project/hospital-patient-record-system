package models

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceStatus string

const (
	InvoicePending InvoiceStatus = "PENDING"
	InvoicePaid    InvoiceStatus = "PAID"
	InvoicePartial InvoiceStatus = "PARTIAL"
)

type Invoice struct {
	ID uint `gorm:"primaryKey"`

	PatientID uint   `gorm:"size:150;index;not null"`
	Patient   Patient `gorm:"foreignKey:PatientID"`

	InvoiceNumber string `gorm:"size:100;uniqueIndex"`

	TotalAmount float64

	PaidAmount float64

	Balance float64

	Status InvoiceStatus `gorm:"default:PENDING"`

	Items []InvoiceItem

	CreatedBy uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}