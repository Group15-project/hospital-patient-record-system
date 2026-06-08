package models

import (
	"time"

	"gorm.io/gorm"
)

type PrescriptionItem struct {
	ID uint `gorm:"primaryKey"`

	PrescriptionID uint         `gorm:"not null;index"`
	Prescription   Prescription `gorm:"foreignKey:PrescriptionID"`

	MedicationID *uint       `gorm:"index"`
	Medication   *Medication `gorm:"foreignKey:MedicationID"`

	Dosage string `gorm:"size:255"`

	Frequency string `gorm:"size:255"`

	Duration string `gorm:"size:255"`

	Quantity int

	Instructions string `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}