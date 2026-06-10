package service

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)

type VitalService interface {
	Create(
		req dto.CreateVitalRequest,
		userID uint,
	) (*dto.VitalResponse, error)

	GetPatientVitals(
		patientID uint,
	) ([]dto.VitalResponse, error)
}

type vitalService struct {
	vitalRepo repository.VitalRepository
}

func NewVitalService(
	vitalRepo repository.VitalRepository,
) VitalService {
	return &vitalService{
		vitalRepo: vitalRepo,
	}
}
func mapVitalResponse(vital *models.Vital) dto.VitalResponse {
	return dto.VitalResponse{
		ID:               vital.ID,
		PatientID:        vital.PatientID,
		Temperature:      vital.Temperature,
		Weight:           vital.Weight,
		Height:           vital.Height,
		PulseRate:        vital.PulseRate,
		BloodSugar:       vital.BloodSugar,
		OxygenSaturation: vital.OxygenSaturation,
		RespiratoryRate:  vital.RespiratoryRate,
		SystolicBP:       vital.SystolicBP,
		DiastolicBP:      vital.DiastolicBP,
		Notes:            vital.Notes,
		RecordedBy:       vital.RecordedBy,
		CreatedAt:        vital.CreatedAt,
	}
}


func (s *vitalService) Create(
	req dto.CreateVitalRequest,
	userID uint,
) (*dto.VitalResponse, error) {

	vital := models.Vital{
		PatientID: req.PatientID,
		ConsultationID: req.ConsultationID,

		Temperature: req.Temperature,
		Weight:      req.Weight,
		Height:      req.Height,

		PulseRate: req.PulseRate,

		BloodSugar: req.BloodSugar,

		OxygenSaturation: req.OxygenSaturation,

		RespiratoryRate: req.RespiratoryRate,

		SystolicBP:  req.SystolicBP,
		DiastolicBP: req.DiastolicBP,

		Notes: req.Notes,

		RecordedBy: userID,
	}

	if err := s.vitalRepo.Create(&vital); err != nil {
		return nil, err
	}

	response := mapVitalResponse(&vital)

	return &response, nil
}




func (s *vitalService) GetPatientVitals(
	patientID uint,
) ([]dto.VitalResponse, error) {

	vitals, err := s.vitalRepo.GetPatientVitals(patientID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.VitalResponse, 0, len(vitals))

	for _, vital := range vitals {
		response = append(
			response,
			mapVitalResponse(&vital),
		)
	}

	return response, nil
}