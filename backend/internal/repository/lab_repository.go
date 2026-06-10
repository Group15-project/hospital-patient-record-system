package repository

import (
	"errors"
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type LabRepository interface {
	CreateRequest(req *models.LabRequest) error

	GetRequestsByPatient(patientID uint) ([]models.LabRequest, error)
	GetResultsByPatient(patientID uint) ([]models.LabResult, error)

	GetPendingRequest() ([]models.LabRequest, error)
	GetRequestByID(id uint) (*models.LabRequest, error)
	UploadResultAndCompleteRequest(
		result *models.LabResult,
	) error
}

type labRepository struct {
	db *gorm.DB
}

func NewLabRepository(db *gorm.DB) LabRepository {
	return &labRepository{
		db: db,
	}
}

func (r *labRepository) CreateRequest(
	req *models.LabRequest,
) error {
	return r.db.Create(req).Error
}

func (r *labRepository) GetRequestByID(
	id uint,
) (*models.LabRequest, error) {

	var req models.LabRequest

	err := r.db.
		Preload("Doctor").
		Preload("Patient").
		First(&req, id).
		Error

	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *labRepository) GetRequestsByPatient(
	patientID uint,
) ([]models.LabRequest, error) {

	var requests []models.LabRequest

	err := r.db.
		Preload("Doctor").
		Preload("Patient").
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&requests).
		Error

	return requests, err
}

func (r *labRepository) GetResultsByPatient(
	patientID uint,
) ([]models.LabResult, error) {

	var results []models.LabResult

	err := r.db.
		Preload("LabRequest").
		Preload("LabTech").
		Joins(
			"JOIN lab_requests ON lab_requests.id = lab_results.lab_request_id",
		).
		Where(
			"lab_requests.patient_id = ?",
			patientID,
		).
		Order("lab_results.created_at DESC").
		Find(&results).
		Error

	return results, err
}

func (r *labRepository) UploadResultAndCompleteRequest(
	result *models.LabResult,
) error {

	return r.db.Transaction(func(tx *gorm.DB) error {

		var request models.LabRequest

		if err := tx.
			First(&request, result.LabRequestID).
			Error; err != nil {
			return err
		}

		// Check if result already exists
		var existing models.LabResult

		err := tx.
			Where("lab_request_id = ?", result.LabRequestID).
			First(&existing).
			Error

		if err == nil {
			return errors.New("lab result already uploaded")
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := tx.Create(result).Error; err != nil {
			return err
		}

		if err := tx.
			Model(&request).
			Update("status", models.LabRequestStatusCompleted).
			Error; err != nil {
			return err
		}

		return nil
	})
}
func (r *labRepository) GetPendingRequest() (
	[]models.LabRequest,
	error,
) {

	var requests []models.LabRequest

	err := r.db.
		Preload("Doctor").
		Preload("Patient").
		Where(
			"status = ?",
			models.LabRequestStatusPending,
		).
		Order("created_at ASC").
		Find(&requests).
		Error

	return requests, err
}
