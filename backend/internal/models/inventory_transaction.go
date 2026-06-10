package models

import (
	"time"

	"gorm.io/gorm"
)

type InventoryTransactionType string

const (
	InventoryStockIn  InventoryTransactionType = "STOCK_IN"
	InventoryStockOut InventoryTransactionType = "STOCK_OUT"
	InventoryAdjustment InventoryTransactionType = "ADJUSTMENT"
)

type InventoryTransaction struct {
	ID uint `gorm:"primaryKey"`

	MedicationID uint       `gorm:"not null"`
	Medication   Medication `gorm:"foreignKey:MedicationID"`

	Type InventoryTransactionType

	Quantity int

	PreviousStock int

	NewStock int

	Reference string

	PerformedBy uint
	User        User `gorm:"foreignKey:PerformedBy"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}