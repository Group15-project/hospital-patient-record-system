package models

import (
	"time"

	"gorm.io/gorm"
)

type Vital struct {
	ID uint `gorm:"primaryKey"`

	PatientID uint    `gorm:"not null;index"`
	Patient   Patient `gorm:"foreignKey:PatientID"`

	ConsultationID *uint
	Consultation   *Consultation `gorm:"foreignKey:ConsultationID"`

	Temperature      float64
	Weight           float64
	Height           float64
	PulseRate        int
	BloodSugar       float64
	OxygenSaturation float64
	RespiratoryRate  int

	SystolicBP  int
	DiastolicBP int

	Notes string `gorm:"type:text"`

	RecordedBy uint
	Recorder   User `gorm:"foreignKey:RecordedBy"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}