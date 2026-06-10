package service

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)

type LabService interface {
	CreateRequest(
		req dto.CreateLabRequestRequest,
		doctorID uint,
	) error

	UploadResult(
		req dto.UploadLabResultRequest,
		labTechID uint,
	) error

	GetPatientRequests(

		patientID uint,
	) ([]models.LabRequest, error)
	GetPatientResults(
		patientID uint,
	) ([]models.LabResult, error)
}

type labService struct {
	repo repository.LabRepository
}

func NewLabService(
	repo repository.LabRepository,
) LabService {
	return &labService{
		repo: repo,
	}
}

func (s *labService) CreateRequest(
	req dto.CreateLabRequestRequest,
	doctorID uint,
) error {

	request := models.LabRequest{
		ConsultationID: req.ConsultationID,
		PatientID:      req.PatientID,
		TestName:       req.TestName,
		ClinicalNote:   req.ClinicalNote,
		RequestedBy:    doctorID,
	}

	return s.repo.CreateRequest(
		&request,
	)
}
func (s *labService) UploadResult(
	req dto.UploadLabResultRequest,
	labTechID uint,
) error {

	result := models.LabResult{
		LabRequestID: req.LabRequestID,
		Result:       req.Result,
		Remarks:      req.Remarks,
		UploadedBy:   labTechID,
	}

	return s.repo.UploadResultAndCompleteRequest(
		&result,
	)
}
func (s *labService) GetPatientRequests(
	patientID uint,
) ([]models.LabRequest, error) {

	return s.repo.GetRequestsByPatient(
		patientID,
	)
}

func (s *labService) GetPatientResults(
	patientID uint,
) ([]models.LabResult, error) {

	return s.repo.GetResultsByPatient(
		patientID,
	)
}