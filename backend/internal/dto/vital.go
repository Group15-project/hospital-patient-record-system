package dto

import "time"

type CreateVitalRequest struct {
	PatientID uint `json:"patient_id" validate:"required"`
	ConsultationID *uint `json:"consultation_id"`
	
	Temperature float64 `json:"temperature"`
	Weight      float64 `json:"weight"`
	Height      float64 `json:"height"`

	PulseRate int `json:"pulse_rate"`

	BloodSugar float64 `json:"blood_sugar"`

	OxygenSaturation float64 `json:"oxygen_saturation"`

	RespiratoryRate int `json:"respiratory_rate"`

	SystolicBP  int `json:"systolic_bp"`
	DiastolicBP int `json:"diastolic_bp"`

	Notes string `json:"notes"`
}

type VitalResponse struct {
	ID               uint      `json:"id"`
	PatientID        uint      `json:"patient_id"`
	Temperature      float64   `json:"temperature"`
	Weight           float64   `json:"weight"`
	Height           float64   `json:"height"`
	PulseRate        int       `json:"pulse_rate"`
	BloodSugar       float64   `json:"blood_sugar"`
	OxygenSaturation float64   `json:"oxygen_saturation"`
	RespiratoryRate  int       `json:"respiratory_rate"`
	SystolicBP       int       `json:"systolic_bp"`
	DiastolicBP      int       `json:"diastolic_bp"`
	Notes            string    `json:"notes"`
	RecordedBy       uint      `json:"recorded_by"`
	CreatedAt        time.Time `json:"created_at"`
}