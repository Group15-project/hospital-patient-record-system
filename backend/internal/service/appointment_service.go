package service

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)

type AppointmentService interface {
	Create(
		req dto.CreateAppointmentRequest,
		userID uint,
	) (*models.Appointment, error)

	GetByPatient(
		patientID uint,
	) ([]models.Appointment, error)

	GetByDoctor(
		doctorID uint,
	) ([]models.Appointment, error)

	UpdateStatus(
		id uint,
		status models.AppointmentStatus,
	) error
	List() ([]models.Appointment, error)
}

type appointmentService struct {
	repo repository.AppointmentRepository
}

func NewAppointmentService(
	repo repository.AppointmentRepository,
) AppointmentService {
	return &appointmentService{
		repo: repo,
	}
}

func (s *appointmentService) Create(
	req dto.CreateAppointmentRequest,
	userID uint,
) (*models.Appointment, error) {

	priority := models.AppointmentNormal

	if req.Priority != "" {
		priority = models.AppointmentPriority(
			req.Priority,
		)
	}

	appointment := models.Appointment{
		PatientID: req.PatientID,
		DoctorID:  req.DoctorID,

		AppointmentDate: req.AppointmentDate,

		Reason: req.Reason,

		Priority: priority,

		CreatedBy: userID,
	}

	err := s.repo.Create(&appointment)

	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (s *appointmentService) GetByPatient(
	patientID uint,
) ([]models.Appointment, error) {

	return s.repo.GetByPatient(
		patientID,
	)
}

func (s *appointmentService) GetByDoctor(
	doctorID uint,
) ([]models.Appointment, error) {

	return s.repo.GetByDoctor(
		doctorID,
	)
}

func (s *appointmentService) UpdateStatus(
	id uint,
	status models.AppointmentStatus,
) error {

	appointment, err := s.repo.GetByID(id)

	if err != nil {
		return err
	}

	appointment.Status = status

	return s.repo.Update(
		appointment,
	)
}

func (s *appointmentService) List() ([]models.Appointment, error) {
	return s.repo.List()
}
