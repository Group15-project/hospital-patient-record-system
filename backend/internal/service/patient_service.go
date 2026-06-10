package service

import (
	"fmt"
	"time"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)

type PatientService interface {
	Create(req dto.CreatePatientRequest, userID uint) (*models.Patient, error)
	GetByID(id uint) (*models.Patient, error)
	List(page, limit int) ([]models.Patient, error)
}

type patientService struct {
	patientRepo repository.PatientRepository
}

func NewPatientService(
	patientRepo repository.PatientRepository,
) PatientService {
	return &patientService{
		patientRepo: patientRepo,
	}
}

func (s *patientService) Create(
	req dto.CreatePatientRequest,
	userID uint,
) (*models.Patient, error) {

	patientNumber := fmt.Sprintf(
		"PAT-%d",
		time.Now().UnixNano(),
	)

	var dob *time.Time

	if req.DateOfBirth != "" {

		parsedDOB, err := time.Parse(
			"2006-01-02",
			req.DateOfBirth,
		)

		if err != nil {
			return nil, err
		}

		dob = &parsedDOB
	}

	patient := models.Patient{
		PatientNumber: patientNumber,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Gender:        req.Gender,
		DateOfBirth:   dob,
		Phone:         req.Phone,
		Email:         req.Email,
		Address:       req.Address,

		EmergencyContactName: req.EmergencyContactName,

		EmergencyContactPhone: req.EmergencyContactPhone,

		BloodGroup: req.BloodGroup,

		CreatedBy: userID,
	}

	err := s.patientRepo.Create(&patient)
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *patientService) GetByID(
	id uint,
) (*models.Patient, error) {
	return s.patientRepo.GetByID(id)
}

// 
func (s *patientService) List(
	page int,
	limit int,
) ([]models.Patient, error) {

	offset := (page - 1) * limit

	return s.patientRepo.List(
		offset,
		limit,
	)
}