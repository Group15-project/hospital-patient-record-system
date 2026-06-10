package models

import (
	"time"

	"gorm.io/gorm"
)

type LabRequestStatus string

const (

	LabRequestStatusPending    LabRequestStatus = "PENDING"
	LabRequestStatusInProgress LabRequestStatus = "IN_PROGRESS"
	LabRequestStatusCompleted  LabRequestStatus = "COMPLETED"

)

type LabRequest struct {
	ID uint `gorm:"primaryKey"`

	ConsultationID uint         `gorm:"not null;index"`
	Consultation   Consultation `gorm:"foreignKey:ConsultationID"`

	PatientID uint    `gorm:"not null;index"`
	Patient   Patient `gorm:"foreignKey:PatientID"`

	RequestedBy uint
	Doctor      User `gorm:"foreignKey:RequestedBy"`

	TestName string `gorm:"size:255;not null"`

	ClinicalNote string `gorm:"type:text"`

	Status LabRequestStatus `gorm:"type:varchar(20);default:'PENDING'"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}