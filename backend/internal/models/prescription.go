package models

import (
	"time"

	"gorm.io/gorm"
)

type PrescriptionStatus string

const (
	PrescriptionPending   PrescriptionStatus = "PENDING"
	PrescriptionDispensed PrescriptionStatus = "DISPENSED"
)

type Prescription struct {
	ID uint `gorm:"primaryKey"`

	ConsultationID uint         `gorm:"not null;index"`
	Consultation   Consultation `gorm:"foreignKey:ConsultationID"`

	PatientID uint    `gorm:"not null;index"`
	Patient   Patient `gorm:"foreignKey:PatientID"`

	DoctorID uint `gorm:"not null"`
	Doctor   User `gorm:"foreignKey:DoctorID"`

	Status PrescriptionStatus `gorm:"type:varchar(20);default:'PENDING';not null"`

	Notes string `gorm:"type:text"`

	Items []PrescriptionItem

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}