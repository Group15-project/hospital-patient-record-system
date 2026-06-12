package service

import (
	"errors"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
	"hospital-backend/internal/utils"

	"time"

	"gorm.io/gorm"
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
	GetDoctors() ([]models.User, error)
	GetProfile(userID uint) (*models.User, error)
	GetStaff() ([]models.User, error)
	DeleteUser(id uint) error
	GetRoles() ([]models.Role, error)
}

type authService struct {
	authRepo     repository.AuthRepository
	auditService AuditService
	jwtSecret    string
}

func NewAuthService(
	authRepo repository.AuthRepository,
	auditService AuditService,
	jwtSecret string,
) AuthService {

	return &authService{
		authRepo:     authRepo,
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
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	IsActive  bool   `json:"is_active"`
}

func (s *authService) Login(
	email string,
	password string,
	auditCtx utils.AuditContext,
) (*LoginResponse, error) {

	user, err := s.authRepo.GetByEmail(email)

	

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
			Phone:     user.Phone,
			Role:      user.Role.Name,
			IsActive:  user.IsActive,
		},
	}, nil
}

func (s *authService) RegisterUser(
	req dto.RegisterUserRequest,
	auditCtx utils.AuditContext,
) error {

	existingUser, err :=
		s.authRepo.GetByEmailIncludingDeleted(
			req.Email,
		)

	if err == nil {

		if existingUser.DeletedAt.Valid {

			hashedPassword, err :=
				utils.HashPassword(
					req.Password,
				)

			if err != nil {
				return err
			}

			existingUser.FirstName =
				req.FirstName

			existingUser.LastName =
				req.LastName

			existingUser.Phone =
				req.Phone

			existingUser.RoleID =
				req.RoleID

			existingUser.HashedPassword =
				hashedPassword

			existingUser.IsActive =
				true

			existingUser.DeletedAt =
				gorm.DeletedAt{}

			return s.authRepo.Update(
				existingUser,
			)
		}

		return errors.New(
			"email already exists",
		)
	}

	hashedPassword, err :=
		utils.HashPassword(
			req.Password,
		)

	if err != nil {
		return err
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,

		HashedPassword:
			hashedPassword,

		RoleID:
			req.RoleID,

		IsActive:
			true,
	}

	if err := s.authRepo.Create(
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

func (s *authService) GetDoctors() (
	[]models.User,
	error,
) {
	return s.authRepo.GetDoctors()
}

func (s *authService) GetProfile(
	userID uint,
) (*models.User, error) {

	return s.authRepo.GetByID(userID)
}

func (s *authService) GetStaff() (
    []models.User,
    error,
) {

    return s.authRepo.GetStaff()
}

func (s *authService) DeleteUser(
    id uint,
) error {

    return s.authRepo.Delete(id)
}

func (s *authService) GetRoles() (
	[]models.Role,
	error,
) {

	return s.authRepo.GetRoles()
}

