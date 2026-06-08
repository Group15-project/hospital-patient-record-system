package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	GetByID(id uint) (*models.Patient, error)
	GetByPatientNumber(number string) (*models.Patient, error)
	List(offset, limit int) ([]models.Patient, error)
	Update(patient *models.Patient) error
	Delete(id uint) error
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{
		db: db,
	}
}

func (r *patientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *patientRepository) GetByID(id uint) (*models.Patient, error) {
	var patient models.Patient

	err := r.db.First(&patient, id).Error

	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (r *patientRepository) GetByPatientNumber(number string) (*models.Patient, error) {
	var patient models.Patient

	err := r.db.
		Where("patient_number = ?", number).
		First(&patient).
		Error

	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (r *patientRepository) List(offset, limit int) ([]models.Patient, error) {
	var patients []models.Patient

	err := r.db.
		Offset(offset).
		Limit(limit).
		Order("id DESC").
		Find(&patients).
		Error

	return patients, err
}

func (r *patientRepository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *patientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}