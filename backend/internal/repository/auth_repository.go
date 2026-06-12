package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	ExistsByEmail(email string) (bool, error)
	GetDoctors() ([]models.User, error)
	FindByID(userID uint) (*models.User, error)
	GetStaff() ([]models.User, error)
	GetRoles() ([]models.Role, error)
	GetByEmailIncludingDeleted(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(
	db *gorm.DB,
) AuthRepository {

	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) GetByID(id uint) (*models.User, error) {
	var user models.User

	err := r.db.
		Preload("Role").
		Preload("Role.Permissions").
		First(&user, id).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.
		Preload("Role").
		Preload("Role.Permissions").
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) Update(
	user *models.User,
) error {

	return r.db.
		Unscoped().
		Save(user).
		Error
}

func (r *authRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *authRepository) ExistsByEmail(
	email string,
) (bool, error) {

	var count int64

	err := r.db.
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).
		Error

	return count > 0, err
}


func (r *authRepository) GetDoctors() (
	[]models.User,
	error,
) {

	var doctors []models.User

	err := r.db.
		Table("users").
		Joins(
			"JOIN roles ON roles.id = users.role_id",
		).
		Where(
			"roles.name = ?",
			models.RoleDoctor,
		).
		Find(&doctors).
		Error

	return doctors, err
}

func (r *authRepository) FindByID(
	userID uint,
) (*models.User, error) {

	var user models.User

	err := r.db.
		Preload("Role").
		First(&user, userID).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *authRepository) GetStaff() (
    []models.User,
    error,
) {

    var users []models.User

    err := r.db.
        Preload("Role").
        Find(&users).
        Error

    return users, err
}

func (r *authRepository) GetRoles() (
	[]models.Role,
	error,
) {

	var roles []models.Role

	err := r.db.
		Order("name ASC").
		Find(&roles).
		Error

	return roles, err
}
func (r *authRepository) GetByEmailIncludingDeleted(
	email string,
) (*models.User, error) {

	var user models.User

	err := r.db.
		Unscoped().
		Preload("Role").
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}