package dto

type CreateLabRequestRequest struct {
	ConsultationID uint `json:"consultation_id" validate:"required"`

	PatientID uint `json:"patient_id" validate:"required"`

	TestName string `json:"test_name" validate:"required"`

	ClinicalNote string `json:"clinical_note"`
}

type UploadLabResultRequest struct {
	LabRequestID uint `json:"lab_request_id" validate:"required"`

	Result string `json:"result" validate:"required"`

	Remarks string `json:"remarks"`
}