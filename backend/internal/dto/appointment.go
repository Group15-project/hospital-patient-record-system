package dto

type CreateAppointmentRequest struct {
	PatientID uint `json:"patient_id" validate:"required"`

	DoctorID uint `json:"doctor_id" validate:"required"`

	AppointmentDate string `json:"appointment_date" validate:"required"`

	Reason string `json:"reason"`
}

type UpdateAppointmentStatusRequest struct {
	Status string `json:"status" validate:"required"`
}