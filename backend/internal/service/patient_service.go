package service

import (
	"encoding/json"
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
	Delete(id uint) error
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
	allergiesJSON, _ := json.Marshal(req.Allergies)

	patient := models.Patient{
		PatientNumber: patientNumber,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Gender:        req.Gender,
		DateOfBirth:   dob,
		Phone:         req.Phone,
		Email:         req.Email,
		Address:       req.Address,
		Allergies: allergiesJSON,

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

func (s *patientService) Delete(
	id uint,
) error {

	_, err := s.patientRepo.GetByID(id)

	if err != nil {
		return err
	}

	return s.patientRepo.Delete(id)
}