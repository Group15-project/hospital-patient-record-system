package  models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID uint `gorm:"primaryKey"`

	PatientNumber string `gorm:"uniqueIndex;not null"`

	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`

	Gender string `gorm:"size:20;not null"`

	DateOfBirth *time.Time

	Phone string `gorm:"size:20"`

	Email string `gorm:"size:150"`
	

	Address string `gorm:"type:text"`

	EmergencyContactName  string `gorm:"size:150"`
	EmergencyContactPhone string `gorm:"size:20"`

	MedicalRecords []MedicalRecord `gorm:"foreignKey:PatientID"`

	BloodGroup string `gorm:"size:10"`

	CreatedBy uint

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}