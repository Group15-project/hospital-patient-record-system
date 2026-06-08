package service

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)

type ConsultationService interface {
	CreateConsultation(
		req dto.CreateConsultationRequest,
		doctorID uint,
	) (*models.Consultation, error)

	AddDiagnosis(
		req dto.AddDiagnosisRequest,
		doctorID uint,
	) error
}

type consultationService struct {
	repo repository.ConsultationRepository
}

func NewConsultationService(
	repo repository.ConsultationRepository,
) ConsultationService {
	return &consultationService{
		repo: repo,
	}
}

func (s *consultationService) CreateConsultation(
	req dto.CreateConsultationRequest,
	doctorID uint,
) (*models.Consultation, error) {

	consultation := models.Consultation{
		PatientID: req.PatientID,
		DoctorID: doctorID,

		ChiefComplaint: req.ChiefComplaint,
	}

	err := s.repo.Create(&consultation)

	if err != nil {
		return nil, err
	}

	return &consultation, nil
}

func (s *consultationService) AddDiagnosis(
	req dto.AddDiagnosisRequest,
	doctorID uint,
) error {

	diagnosis := models.Diagnosis{
		ConsultationID: req.ConsultationID,

		PrimaryDiagnosis: req.Diagnosis,

		TreatmentPlan: req.TreatmentPlan,

		Notes: req.Notes,

		CreatedBy: doctorID,
	}

	return s.repo.CreateDiagnosis(
		&diagnosis,
	)
}