package models

import (
	"time"

	"gorm.io/gorm"
)

type AppointmentStatus string
type AppointmentPriority string

const (
	AppointmentScheduled AppointmentStatus = "SCHEDULED"
	AppointmentCompleted AppointmentStatus = "COMPLETED"
	AppointmentCancelled AppointmentStatus = "CANCELLED"
	AppointmentNoShow    AppointmentStatus = "NO_SHOW"
	AppointmentNormal    AppointmentPriority = "NORMAL"

	AppointmentUrgent    AppointmentPriority = "URGENT"
	AppointmentEmergency AppointmentPriority = "EMERGENCY"


)

type Appointment struct {
	ID uint `gorm:"primaryKey"`

	PatientID uint    `gorm:"not null;index"`
	Patient   Patient `gorm:"foreignKey:PatientID"`

	DoctorID uint `gorm:"not null;index"`
	Doctor   User `gorm:"foreignKey:DoctorID"`

	AppointmentDate time.Time `gorm:"not null"`

	Reason string `gorm:"type:text"`
	Priority AppointmentPriority `gorm:"default:NORMAL"`

	Status AppointmentStatus `gorm:"default:SCHEDULED"`

	CreatedBy uint
	Creator   User `gorm:"foreignKey:CreatedBy"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}