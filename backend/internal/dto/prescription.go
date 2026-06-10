package dto

type PrescriptionItemRequest struct {
MedicationID uint `json:"medication_id" validate:"required"`

	Dosage string `json:"dosage"`

	Frequency string `json:"frequency"`

	Duration string `json:"duration"`

	Quantity int `json:"quantity"`

	Instructions string `json:"instructions"`
}

type CreatePrescriptionRequest struct {
	ConsultationID uint `json:"consultation_id" validate:"required"`
	
	PatientID uint `json:"patient_id" validate:"required"`

	Notes string `json:"notes"`

	Items []PrescriptionItemRequest `json:"items" validate:"required,min=1"`
}