package service

import (
	"fmt"
	"time"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)

type BillingService interface {
	CreateInvoice(
		req dto.CreateInvoiceRequest,
		userID uint,
	) error



	GetOutstandingBalance(
		patientID uint,
	) (float64, error)

	RecordPayment(
		req dto.CreatePaymentRequest,
		userID uint,
	) error

	GetPatientInvoices(
		patientID uint,
	) ([]models.Invoice, error)
}

type billingService struct {
	repo repository.BillingRepository
}

func NewBillingService(
	repo repository.BillingRepository,
) BillingService {
	return &billingService{
		repo: repo,
	}
}

func (s *billingService) CreateInvoice(
	req dto.CreateInvoiceRequest,
	userID uint,
) error {

	if len(req.Items) == 0 {
		return fmt.Errorf("invoice must contain at least one item")
	}

	invoice := models.Invoice{
		PatientID: req.PatientID,
		CreatedBy: userID,
		InvoiceNumber: fmt.Sprintf(
			"INV-%d",
			time.Now().UnixNano(),
		),
		Status: models.InvoicePending,
	}

	total := 0.0

	for _, item := range req.Items {

		amount := float64(item.Quantity) *
			item.UnitPrice

		total += amount

		invoice.Items = append(
			invoice.Items,
			models.InvoiceItem{
				Description: item.Description,
				Quantity:    item.Quantity,
				UnitPrice:   item.UnitPrice,
				Amount:      amount,
			},
		)
	}

	invoice.TotalAmount = total
	invoice.PaidAmount = 0
	invoice.Balance = total

	return s.repo.CreateInvoice(
		&invoice,
	)
}

func (s *billingService) RecordPayment(
	req dto.CreatePaymentRequest,
	userID uint,
) error {

	invoice, err := s.repo.GetInvoice(
		req.InvoiceID,
	)

	if err != nil {
		return err
	}

	if req.Amount <= 0 {
		return fmt.Errorf(
			"payment amount must be greater than zero",
		)
	}

	if req.Amount > invoice.Balance {
		return fmt.Errorf(
			"payment exceeds outstanding balance",
		)
	}

	payment := models.Payment{
		InvoiceID: req.InvoiceID,
		Amount:    req.Amount,
		Method: models.PaymentMethod(
			req.Method,
		),
		Reference:  req.Reference,
		ReceivedBy: userID,
	}

	if err := s.repo.CreatePayment(
		&payment,
	); err != nil {
		return err
	}

	invoice.PaidAmount += req.Amount
	invoice.Balance =
		invoice.TotalAmount - invoice.PaidAmount

	switch {
	case invoice.Balance <= 0:
		invoice.Status = models.InvoicePaid

	case invoice.PaidAmount > 0:
		invoice.Status = models.InvoicePartial

	default:
		invoice.Status = models.InvoicePending
	}

	return s.repo.UpdateInvoice(
		invoice,
	)
}

func (s *billingService) GetPatientInvoices(
	patientID uint,
) ([]models.Invoice, error) {

	return s.repo.GetPatientInvoices(
		patientID,
	)
}


func (s *billingService) GetOutstandingBalance(
	patientID uint,
) (float64, error) {

	return s.repo.GetOutstandingBalance(
		patientID,
	)
}