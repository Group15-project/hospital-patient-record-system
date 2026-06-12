package dto

import "time"

type CreateAppointmentRequest struct {
	PatientID uint `json:"patient_id" validate:"required"`

	DoctorID uint `json:"doctor_id" validate:"required"`

	AppointmentDate time.Time `json:"appointment_date" validate:"required"`

	Reason string `json:"reason"`
	Priority string `json:"priority"`

}

type UpdateAppointmentStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

type RescheduleAppointmentRequest struct {
	AppointmentDate string `json:"appointment_date" binding:"required"`
}