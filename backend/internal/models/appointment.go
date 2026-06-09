package models

import (
	"time"

	"gorm.io/gorm"
)

type AppointmentStatus string

const (
	AppointmentScheduled AppointmentStatus = "SCHEDULED"
	AppointmentCompleted AppointmentStatus = "COMPLETED"
	AppointmentCancelled AppointmentStatus = "CANCELLED"
	AppointmentNoShow    AppointmentStatus = "NO_SHOW"
)

type Appointment struct {
	ID uint `gorm:"primaryKey"`

	PatientID uint    `gorm:"not null;index"`
	Patient   Patient `gorm:"foreignKey:PatientID"`

	DoctorID uint `gorm:"not null;index"`
	Doctor   User `gorm:"foreignKey:DoctorID"`

	AppointmentDate time.Time `gorm:"not null"`

	Reason string `gorm:"type:text"`

	Status AppointmentStatus `gorm:"default:SCHEDULED"`

	CreatedBy uint
	Creator   User `gorm:"foreignKey:CreatedBy"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}