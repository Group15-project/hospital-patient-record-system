package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)




type MedicationRepository interface {

	Create(*models.Medication) error
	GetByID(uint) (*models.Medication, error)
	List() ([]models.Medication, error)
	Update(*models.Medication) error
	CreateTransaction(
		*models.InventoryTransaction,
	) error

}

type medicationRepository struct {
	db *gorm.DB
}

func NewMedicationRepository(
	db *gorm.DB,
) MedicationRepository {
	return &medicationRepository{
		db: db,
	}
}

func (r *medicationRepository) Create(
	medication *models.Medication,
) error {
	return r.db.Create(medication).Error
}

func (r *medicationRepository) GetByID(
	id uint,
) (*models.Medication, error) {

	var medication models.Medication

	err := r.db.First(
		&medication,
		id,
	).Error

	if err != nil {
		return nil, err
	}

	return &medication, nil
}

func (r *medicationRepository) List() (
	[]models.Medication,
	error,
) {

	var medications []models.Medication

	err := r.db.
		Order("name ASC").
		Find(&medications).
		Error

	return medications, err
}

func (r *medicationRepository) Update(
	medication *models.Medication,
) error {
	return r.db.Save(medication).Error
}

func (r *medicationRepository) CreateTransaction(
	tx *models.InventoryTransaction,
) error {
	return r.db.Create(tx).Error
}