package  dto

type CreatePatientRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`

	Gender string `json:"gender" validate:"required"`

	DateOfBirth string `json:"date_of_birth"`

	Phone string `json:"phone"`

	Email string `json:"email"`

	Address string `json:"address"`

	EmergencyContactName string `json:"emergency_contact_name"`

	EmergencyContactPhone string `json:"emergency_contact_phone"`

	BloodGroup string `json:"blood_group"`
}

type UpdatePatientRequest struct {
	Phone string `json:"phone"`

	Email string `json:"email"`

	Address string `json:"address"`

	EmergencyContactName string `json:"emergency_contact_name"`

	EmergencyContactPhone string `json:"emergency_contact_phone"`

	BloodGroup string `json:"blood_group"`

	IsActive bool `json:"is_active"`
}