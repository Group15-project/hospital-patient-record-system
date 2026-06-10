package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type MedicalDocumentRepository interface {
	Create(*models.MedicalDocument) error

	GetByID(uint) (*models.MedicalDocument, error)

	GetByPatient(
		patientID uint,
	) ([]models.MedicalDocument, error)
}

type medicalDocumentRepository struct {
	db *gorm.DB
}

func NewMedicalDocumentRepository(
	db *gorm.DB,
) MedicalDocumentRepository {
	return &medicalDocumentRepository{
		db: db,
	}
}

func (r *medicalDocumentRepository) Create(
	doc *models.MedicalDocument,
) error {
	return r.db.Create(doc).Error
}

func (r *medicalDocumentRepository) GetByID(
	id uint,
) (*models.MedicalDocument, error) {

	var doc models.MedicalDocument

	err := r.db.
		Preload("Uploader").
		First(&doc, id).
		Error

	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (r *medicalDocumentRepository) GetByPatient(
	patientID uint,
) ([]models.MedicalDocument, error) {

	var docs []models.MedicalDocument

	err := r.db.
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&docs).
		Error

	return docs, err
}