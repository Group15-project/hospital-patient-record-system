package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)



type AppointmentRepository interface {

	Create(*models.Appointment) error
	GetByID(uint) (*models.Appointment, error)
	GetByPatient(uint) ([]models.Appointment, error)
	GetByDoctor(uint) ([]models.Appointment, error)
	Update(*models.Appointment) error
	List() ([]models.Appointment, error)
	CountEmergencyAppointments() (
	int64,
	error,
)
}

type appointmentRepository struct {

	db *gorm.DB

}

func NewAppointmentRepository(

	db *gorm.DB,

) AppointmentRepository {

	return &appointmentRepository{
		db: db,
	}

}

func (r *appointmentRepository) Create(
	appointment *models.Appointment,
) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) GetByID(
	id uint,
) (*models.Appointment, error) {

	var appointment models.Appointment

	err := r.db.
		Preload("Patient").
		Preload("Doctor").
		First(&appointment, id).
		Error

	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (r *appointmentRepository) GetByPatient(
	patientID uint,
) ([]models.Appointment, error) {

	var appointments []models.Appointment

	err := r.db.
		Preload("Doctor").
		Where("patient_id = ?", patientID).
		Order("appointment_date DESC").
		Find(&appointments).
		Error

	return appointments, err
}

func (r *appointmentRepository) GetByDoctor(
	doctorID uint,
) ([]models.Appointment, error) {

	var appointments []models.Appointment

	err := r.db.
		Preload("Patient").
		Where("doctor_id = ?", doctorID).
		Order("appointment_date DESC").
		Find(&appointments).
		Error

	return appointments, err
}

func (r *appointmentRepository) Update(
	appointment *models.Appointment,
) error {
	return r.db.Save(appointment).Error
}

func (r *appointmentRepository) List() ([]models.Appointment, error) {

	var appointments []models.Appointment

	err := r.db.
		Preload("Patient").
		Preload("Doctor").
		Order("appointment_date DESC").
		Find(&appointments).
		Error

	return appointments, err
}

func (r *appointmentRepository) CountEmergencyAppointments() (
	int64,
	error,
) {

	var count int64

	err := r.db.
		Model(&models.Appointment{}).
		Where(
			"priority = ?",
			models.AppointmentEmergency,
		).
		Where(
			"status = ?",
			models.AppointmentScheduled,
		).
		Count(&count).
		Error

	return count, err
}