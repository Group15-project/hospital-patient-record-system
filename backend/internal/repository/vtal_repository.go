package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type VitalRepository interface {
	Create(vital *models.Vital) error
	GetByID(id uint) (*models.Vital, error)
	GetPatientVitals(patientID uint) ([]models.Vital, error)
}

type vitalRepository struct {
	db *gorm.DB
}

func NewVitalRepository(db *gorm.DB) VitalRepository {
	return &vitalRepository{
		db: db,
	}
}

func (r *vitalRepository) Create(
	vital *models.Vital,
) error {
	return r.db.Create(vital).Error
}

func (r *vitalRepository) GetByID(
	id uint,
) (*models.Vital, error) {

	var vital models.Vital

	err := r.db.
		Preload("Patient").
		Preload("Recorder").
		First(&vital, id).
		Error

	if err != nil {
		return nil, err
	}

	return &vital, nil
}

func (r *vitalRepository) GetPatientVitals(
	patientID uint,
) ([]models.Vital, error) {

	var vitals []models.Vital

	err := r.db.
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&vitals).
		Error

	return vitals, err
}
