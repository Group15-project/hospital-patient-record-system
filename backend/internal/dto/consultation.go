package dto

type CreateConsultationRequest struct {
	PatientID uint `json:"patient_id" validate:"required"`

	ChiefComplaint string `json:"chief_complaint"`
}

type AddDiagnosisRequest struct {
	ConsultationID uint `json:"consultation_id" validate:"required"`

	Diagnosis string `json:"diagnosis" validate:"required"`

	TreatmentPlan string `json:"treatment_plan"`

	Notes string `json:"notes"`
}