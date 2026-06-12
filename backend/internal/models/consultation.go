package models

import (
	"time"

	"gorm.io/gorm"
)

type ConsultationStatus string

const (
	ConsultationOpen   ConsultationStatus = "OPEN"
	ConsultationClosed ConsultationStatus = "CLOSED"
)

type Consultation struct {
	ID uint `gorm:"primaryKey"`

	PatientID uint    `gorm:"not null;index"`
	Patient   Patient `gorm:"foreignKey:PatientID"`

	DoctorID uint `gorm:"not null;index"`
	Doctor   User `gorm:"foreignKey:DoctorID"`

	ChiefComplaint string `gorm:"type:text"`

	Status ConsultationStatus `gorm:"type:varchar(20);default:'OPEN';not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}