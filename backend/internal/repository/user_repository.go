package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	ExistsByEmail(email string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
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

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
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

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) ExistsByEmail(
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