package dto

type UploadMedicalDocumentRequest struct {
	PatientID uint `form:"patient_id" binding:"required"`

	ConsultationID *uint `form:"consultation_id"`

	Title string `form:"title"`

	DocumentType string `form:"document_type"`
}