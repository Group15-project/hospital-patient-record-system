package models

import (
	"time"

	"gorm.io/gorm"
)

type MedicalRecord struct {
	ID uint `gorm:"primaryKey" json:"id"`

	PatientID uint `gorm:"not null;index" json:"patientId"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"patient"`

	Title string `gorm:"size:255;not null" json:"title"`

	Description string `gorm:"type:text" json:"description"`

	Type string `gorm:"size:100" json:"type"`

	Severity string `gorm:"size:50" json:"severity"`

	DoctorName string `gorm:"size:255" json:"doctorName"`

	Prescription string `gorm:"type:text" json:"prescription"`

	RecordDate time.Time `json:"date"`

	CreatedBy uint `json:"createdBy"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}