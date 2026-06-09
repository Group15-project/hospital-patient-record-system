package service

import (
	"errors"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
	"hospital-backend/internal/utils"

	"time"
)

type AuthService interface {
	Login(
		email string,
		password string,
		auditCtx utils.AuditContext,
	) (*LoginResponse, error)

	RegisterUser(

		req dto.RegisterUserRequest,
		auditCtx utils.AuditContext,
	) error
}

type authService struct {
	userRepo     repository.UserRepository
	auditService AuditService
	jwtSecret    string
}

func NewAuthService(
	userRepo repository.UserRepository,
	auditService AuditService,
	jwtSecret string,
) AuthService {

	return &authService{
		userRepo:     userRepo,
		auditService: auditService,
		jwtSecret:    jwtSecret,
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
	auditCtx utils.AuditContext,
) (*LoginResponse, error) {

	user, err := s.userRepo.GetByEmail(email)

	if err != nil {

		s.auditService.Log(
			auditCtx,
			"LOGIN_FAILED",
			"AUTH",
			"",
			"user not found",
		)

		return nil, errors.New("invalid credentials")
	}

	if err != nil {

		s.auditService.Log(
			auditCtx,
			"LOGIN_FAILED",
			"AUTH",
			"",
			"user not found",
		)

		return nil, errors.New("invalid credentials")
	}

	if err := utils.ComparePassword(
		user.HashedPassword,
		password,
	); err != nil {

		auditCtx.UserID = &user.ID

		s.auditService.Log(
			auditCtx,
			"LOGIN_FAILED",
			"AUTH",
			"",
			"invalid password",
		)

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
	auditCtx.UserID = &user.ID

	s.auditService.Log(
		auditCtx,
		"LOGIN_SUCCESS",
		"AUTH",
		"",
		"user logged in",
	)

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

func (s *authService) RegisterUser(
	req dto.RegisterUserRequest,
	auditCtx utils.AuditContext,
) error {
	exists, err := s.userRepo.ExistsByEmail(
		req.Email,
	)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("email already exists")
	}
	hashedPassword, err := utils.HashPassword(
		req.Password,
	)

	if err != nil {
		return err
	}
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,

		Email: req.Email,
		Phone: req.Phone,

		HashedPassword: hashedPassword,

		RoleID: req.RoleID,

		IsActive: true,
	}
	if err := s.userRepo.Create(
		&user,
	); err != nil {
		return err
	}
	s.auditService.Log(
		auditCtx,
		"CREATE",
		"USER",
		"",
		"staff account created",
	)
	return nil
}
