package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type PrescriptionRepository interface {
	Create(
		prescription *models.Prescription,
	) error

	GetByID(
		id uint,
	) (*models.Prescription, error)

	GetByPatient(
		patientID uint,
	) ([]models.Prescription, error)
}
type prescriptionRepository struct {
	db *gorm.DB
}

func NewPrescriptionRepository(
	db *gorm.DB,
) PrescriptionRepository {
	return &prescriptionRepository{
		db: db,
	}
}

func (r *prescriptionRepository) Create(
	prescription *models.Prescription,
) error {

	return r.db.Create(
		prescription,
	).Error
}
func (r *prescriptionRepository) GetByID(
	id uint,
) (*models.Prescription, error) {

	var prescription models.Prescription

	err := r.db.
		Preload("Items").
		Preload("Items.Medication").
		Preload("Doctor").
		Preload("Patient").
		First(&prescription, id).
		Error

	if err != nil {
		return nil, err
	}

	return &prescription, nil
}

func (r *prescriptionRepository) GetByPatient(
	patientID uint,
) ([]models.Prescription, error) {

	var prescriptions []models.Prescription

	err := r.db.
		Preload("Items").
		Preload("Items.Medication").
		Preload("Doctor").
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&prescriptions).
		Error

	return prescriptions, err
}