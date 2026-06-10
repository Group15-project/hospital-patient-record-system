package dto


type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}



type RegisterUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`

	Email string `json:"email" validate:"required,email"`

	Phone string `json:"phone"`

	Password string `json:"password" validate:"required,min=8"`

	RoleID uint `json:"role_id" validate:"required"`
}

type ProfileResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}