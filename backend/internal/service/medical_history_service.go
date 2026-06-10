package service

import (
	"time"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)

type MedicalRecordService interface {
	Create(
		req dto.CreateMedicalRecordRequest,
		userID uint,
	) (*models.MedicalRecord, error)

	GetByID(
		id uint,
	) (*models.MedicalRecord, error)

	GetByPatientID(
		patientID uint,
	) ([]models.MedicalRecord, error)

	List() (
		[]models.MedicalRecord,
		error,
	)

	Update(
		id uint,
		req dto.UpdateMedicalRecordRequest,
	) (*models.MedicalRecord, error)

	Delete(
		id uint,
	) error
}

type medicalRecordService struct {
	repo repository.MedicalRecordRepository
}

func NewMedicalRecordService(
	repo repository.MedicalRecordRepository,
) MedicalRecordService {
	return &medicalRecordService{
		repo: repo,
	}
}

func (s *medicalRecordService) Create(
	req dto.CreateMedicalRecordRequest,
	userID uint,
) (*models.MedicalRecord, error) {

	recordDate := time.Now()

	if req.Date != "" {
		parsedDate, err := time.Parse(
			"2006-01-02",
			req.Date,
		)

		if err != nil {
			return nil, err
		}

		recordDate = parsedDate
	}

	record := models.MedicalRecord{
		PatientID:    req.PatientID,
		Title:        req.Title,
		Description:  req.Description,
		Type:         req.Type,
		Severity:     req.Severity,
		DoctorName:   req.DoctorName,
		Prescription: req.Prescription,
		RecordDate:   recordDate,
		CreatedBy:    userID,
	}

	err := s.repo.Create(&record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (s *medicalRecordService) GetByID(
	id uint,
) (*models.MedicalRecord, error) {
	return s.repo.GetByID(id)
}

func (s *medicalRecordService) GetByPatientID(
	patientID uint,
) ([]models.MedicalRecord, error) {
	return s.repo.GetByPatientID(patientID)
}

func (s *medicalRecordService) List() (
	[]models.MedicalRecord,
	error,
) {
	return s.repo.List()
}

func (s *medicalRecordService) Update(
	id uint,
	req dto.UpdateMedicalRecordRequest,
) (*models.MedicalRecord, error) {

	record, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		record.Title = req.Title
	}

	if req.Description != "" {
		record.Description = req.Description
	}

	if req.Type != "" {
		record.Type = req.Type
	}

	if req.Severity != "" {
		record.Severity = req.Severity
	}

	if req.DoctorName != "" {
		record.DoctorName = req.DoctorName
	}

	if req.Prescription != "" {
		record.Prescription = req.Prescription
	}

	if req.Date != "" {

		parsedDate, err := time.Parse(
			"2006-01-02",
			req.Date,
		)

		if err != nil {
			return nil, err
		}

		record.RecordDate = parsedDate
	}

	err = s.repo.Update(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (s *medicalRecordService) Delete(
	id uint,
) error {
	return s.repo.Delete(id)
}