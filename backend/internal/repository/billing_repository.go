package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type BillingRepository interface {
	CreateInvoice(*models.Invoice) error

	GetInvoice(uint) (*models.Invoice, error)

	CreatePayment(*models.Payment) error

	UpdateInvoice(*models.Invoice) error
GetOutstandingBalance(
	patientID uint,
) (float64, error)

	GetPatientInvoices(
		patientID uint,
	) ([]models.Invoice, error)
}

type billingRepository struct {
	db *gorm.DB
}

func NewBillingRepository(
	db *gorm.DB,
) BillingRepository {
	return &billingRepository{
		db: db,
	}
}

func (r *billingRepository) CreateInvoice(
	invoice *models.Invoice,
) error {
	return r.db.Create(invoice).Error
}

func (r *billingRepository) GetInvoice(
	id uint,
) (*models.Invoice, error) {

	var invoice models.Invoice

	err := r.db.
		Preload("Items").
		First(&invoice, id).
		Error

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (r *billingRepository) UpdateInvoice(
	invoice *models.Invoice,
) error {
	return r.db.Save(invoice).Error
}

func (r *billingRepository) GetPatientInvoices(
	patientID uint,
) ([]models.Invoice, error) {

	var invoices []models.Invoice

	err := r.db.
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&invoices).
		Error

	return invoices, err
}

func (r *billingRepository) CreatePayment(
	payment *models.Payment,
) error {
	return r.db.Create(payment).Error
}

func (r *billingRepository) GetOutstandingBalance(
	patientID uint,
) (float64, error) {

	var total float64

	err := r.db.
		Model(&models.Invoice{}).
		Where(
			"patient_id = ?",
			patientID,
		).
		Select("COALESCE(SUM(balance),0)").
		Scan(&total).
		Error

	return total, err
}