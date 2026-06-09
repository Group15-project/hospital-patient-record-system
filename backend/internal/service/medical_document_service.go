package service

import (
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)

type MedicalDocumentService interface {
	Upload(
		document *models.MedicalDocument,
	) error

	GetPatientDocuments(
		patientID uint,
	) ([]models.MedicalDocument, error)
}

type medicalDocumentService struct {
	repo repository.MedicalDocumentRepository
}

func NewMedicalDocumentService(
	repo repository.MedicalDocumentRepository,
) MedicalDocumentService {
	return &medicalDocumentService{
		repo: repo,
	}
}

func (s *medicalDocumentService) Upload(
	document *models.MedicalDocument,
) error {

	return s.repo.Create(
		document,
	)
}

func (s *medicalDocumentService) GetPatientDocuments(
	patientID uint,
) ([]models.MedicalDocument, error) {

	return s.repo.GetByPatient(
		patientID,
	)
}