package service

import (
	"errors"

	"hospital-backend/internal/repository"
	"hospital-backend/internal/utils"

	"time"
)

type AuthService interface {
	Login(email string, password string) (*LoginResponse, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthService(

	userRepo repository.UserRepository,
	jwtSecret string,

) AuthService {

	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}

}

type LoginResponse struct {
	AccessToken string `json:"access_token"`

	User UserResponse `json:"user"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

func (s *authService) Login(
	email string,
	password string,
) (*LoginResponse, error) {

	user, err := s.userRepo.GetByEmail(email)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("account disabled")
	}

	if err := utils.ComparePassword(
	user.HashedPassword,
	password,
); err != nil {
	return nil, errors.New("invalid credentials")
}

	token, err := utils.GenerateAccessToken(
		user.ID,
		user.Role.Name,
		s.jwtSecret,
		24*time.Hour,
	)

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken: token,
		User: UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role.Name,
		},
	}, nil
}
