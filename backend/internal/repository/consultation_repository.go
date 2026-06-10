package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type ConsultationRepository interface {
	Create(consultation *models.Consultation) error
	GetByID(id uint) (*models.Consultation, error)

	CreateDiagnosis(
		diagnosis *models.Diagnosis,
	) error
	
		GetByPatient(
		patientID uint,
	) ([]models.Consultation, error)
	
	GetDiagnoses(
		consultationID uint,
	) ([]models.Diagnosis, error)
}

type consultationRepository struct {
	db *gorm.DB
}

func NewConsultationRepository(
	db *gorm.DB,
) ConsultationRepository {
	return &consultationRepository{
		db: db,
	}
}

func (r *consultationRepository) Create(
	consultation *models.Consultation,
) error {
	return r.db.Create(consultation).Error
}

func (r *consultationRepository) GetByID(
	id uint,
) (*models.Consultation, error) {

	var consultation models.Consultation

	err := r.db.
		Preload("Patient").
		Preload("Doctor").
		First(&consultation, id).
		Error

	if err != nil {
		return nil, err
	}

	return &consultation, nil
}

func (r *consultationRepository) CreateDiagnosis(
	diagnosis *models.Diagnosis,
) error {

	return r.db.Create(diagnosis).Error
}
func (r *consultationRepository) GetDiagnoses(
	consultationID uint,
) ([]models.Diagnosis, error) {

	var diagnoses []models.Diagnosis

	err := r.db.
		Where(
			"consultation_id = ?",
			consultationID,
		).
		Order("created_at DESC").
		Find(&diagnoses).
		Error

	return diagnoses, err
}

func (r *consultationRepository) GetByPatient(
	patientID uint,
) ([]models.Consultation, error) {

	var consultations []models.Consultation

	err := r.db.
		Preload("Doctor").
		Where(
			"patient_id = ?",
			patientID,
		).
		Order("created_at DESC").
		Find(&consultations).
		Error

	return consultations, err
}