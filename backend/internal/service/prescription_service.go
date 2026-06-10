package service

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)




type PrescriptionService interface {

	Create(
		req dto.CreatePrescriptionRequest,
		doctorID uint,
	) error
	GetByPatient(
		patientID uint,
	) ([]models.Prescription, error)

}

type prescriptionService struct {
	repo repository.PrescriptionRepository
}

func NewPrescriptionService(
	repo repository.PrescriptionRepository,
) PrescriptionService {
	return &prescriptionService{
		repo: repo,
	}
}

func (s *prescriptionService) Create(
	req dto.CreatePrescriptionRequest,
	doctorID uint,
) error {

	prescription := models.Prescription{
		ConsultationID: req.ConsultationID,
		PatientID:      req.PatientID,
		DoctorID:       doctorID,
		Notes:          req.Notes,
	}

	for _, item := range req.Items {

		prescription.Items = append(
			prescription.Items,
			models.PrescriptionItem{
				MedicationID: &item.MedicationID,
				Dosage: item.Dosage,
				Frequency: item.Frequency,
				Duration: item.Duration,
				Quantity: item.Quantity,
				Instructions: item.Instructions,
			},
		)
	}

	return s.repo.Create(
		&prescription,
	)
}

func (s *prescriptionService) GetByPatient(
	patientID uint,
) ([]models.Prescription, error) {

	return s.repo.GetByPatient(
		patientID,
	)
}