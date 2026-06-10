package dto

type CreateMedicalRecordRequest struct {
	PatientID uint `json:"patientId" validate:"required"`

	Title string `json:"title" validate:"required"`

	Description string `json:"description" validate:"required"`

	Type string `json:"type"`

	Severity string `json:"severity"`

	DoctorName string `json:"doctorName"`

	Prescription string `json:"prescription"`

	Date string `json:"date" validate:"required"`
}

type UpdateMedicalRecordRequest struct {
	Title string `json:"title"`

	Description string `json:"description"`

	Type string `json:"type"`

	Severity string `json:"severity"`

	DoctorName string `json:"doctorName"`

	Prescription string `json:"prescription"`

	Date string `json:"date"`
}