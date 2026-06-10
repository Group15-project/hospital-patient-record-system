package models

import (
	"time"

	"gorm.io/gorm"
)

type Diagnosis struct {
	ID uint `gorm:"primaryKey"`

	ConsultationID uint         `gorm:"not null;index"`
	Consultation   Consultation `gorm:"foreignKey:ConsultationID"`

	PrimaryDiagnosis   string `gorm:"type:text;not null"`
	SecondaryDiagnosis string
	ICD10Code          string
	TreatmentPlan      string `gorm:"type:text"`

	Notes string `gorm:"type:text"`

	CreatedBy uint
	Doctor    User `gorm:"foreignKey:CreatedBy"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
