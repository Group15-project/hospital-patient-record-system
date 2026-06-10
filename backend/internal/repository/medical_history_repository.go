package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type MedicalRecordRepository interface {
	Create(record *models.MedicalRecord) error
	GetByID(id uint) (*models.MedicalRecord, error)
	GetByPatientID(patientID uint) ([]models.MedicalRecord, error)
	List() ([]models.MedicalRecord, error)
	Update(record *models.MedicalRecord) error
	Delete(id uint) error
}

type medicalRecordRepository struct {
	db *gorm.DB
}

func NewMedicalRecordRepository(
	db *gorm.DB,
) MedicalRecordRepository {
	return &medicalRecordRepository{
		db: db,
	}
}

func (r *medicalRecordRepository) Create(
	record *models.MedicalRecord,
) error {

	err := r.db.Create(record).Error
	if err != nil {
		return err
	}

	return r.db.
		Preload("Patient").
		First(record, record.ID).
		Error
}

func (r *medicalRecordRepository) GetByID(
	id uint,
) (*models.MedicalRecord, error) {

	var record models.MedicalRecord

	err := r.db.
		Preload("Patient").
		First(&record, id).
		Error

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *medicalRecordRepository) GetByPatientID(
	patientID uint,
) ([]models.MedicalRecord, error) {

	var records []models.MedicalRecord

	err := r.db.
		Where("patient_id = ?", patientID).
		Order("record_date DESC").
		Find(&records).
		Error

	return records, err
}

func (r *medicalRecordRepository) List() (
	[]models.MedicalRecord,
	error,
) {

	var records []models.MedicalRecord

	err := r.db.
		Preload("Patient").
		Order("record_date DESC").
		Find(&records).
		Error

	return records, err
}

func (r *medicalRecordRepository) Update(
	record *models.MedicalRecord,
) error {
	return r.db.Save(record).Error
}

func (r *medicalRecordRepository) Delete(
	id uint,
) error {
	return r.db.Delete(
		&models.MedicalRecord{},
		id,
	).Error
}
