package models

import (
	"time"

	"gorm.io/gorm"
)

type MedicalDocumentType string

const (
	DocumentClinicalNote MedicalDocumentType = "CLINICAL_NOTE"
	DocumentXRay         MedicalDocumentType = "XRAY"
	DocumentMRI          MedicalDocumentType = "MRI"
	DocumentCTScan       MedicalDocumentType = "CT_SCAN"
	DocumentUltrasound   MedicalDocumentType = "ULTRASOUND"
	DocumentLabReport    MedicalDocumentType = "LAB_REPORT"
	DocumentPrescription MedicalDocumentType = "PRESCRIPTION"
	DocumentOther        MedicalDocumentType = "OTHER"
)

type MedicalDocument struct {
	ID uint `gorm:"primaryKey"`

	PatientID uint    `gorm:"not null;index"`
	Patient   Patient `gorm:"foreignKey:PatientID"`

	ConsultationID *uint
	Consultation   *Consultation

	Title string `gorm:"size:255"`

	DocumentType MedicalDocumentType `gorm:"type:varchar(50);not null"`

	FileName string

	FilePath string

	MimeType string

	FileSize int64

	UploadedBy uint
	Uploader  User `gorm:"foreignKey:UploadedBy"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}