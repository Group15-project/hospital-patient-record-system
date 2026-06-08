package service

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
)




type MedicationService interface {

	Create(
		req dto.CreateMedicationRequest,
		userID uint,
	) error
	List() ([]models.Medication, error)
	AddStock(
		req dto.StockAdjustmentRequest,
		userID uint,
	) error

}

type medicationService struct {
	repo repository.MedicationRepository
}

func NewMedicationService(
	repo repository.MedicationRepository,
) MedicationService {
	return &medicationService{
		repo: repo,
	}
}

func (s *medicationService) Create(
	req dto.CreateMedicationRequest,
	userID uint,
) error {

	med := models.Medication{
		Name: req.Name,
		GenericName: req.GenericName,
		Strength: req.Strength,
		Form: req.Form,
		UnitPrice: req.UnitPrice,
		QuantityInStock: req.QuantityInStock,
		ReorderLevel: req.ReorderLevel,
	}

	return s.repo.Create(&med)
}

func (s *medicationService) List() (
	[]models.Medication,
	error,
) {
	return s.repo.List()
}

func (s *medicationService) AddStock(
	req dto.StockAdjustmentRequest,
	userID uint,
) error {

	med, err := s.repo.GetByID(
		req.MedicationID,
	)

	if err != nil {
		return err
	}

	previous := med.QuantityInStock

	med.QuantityInStock += req.Quantity

	if err := s.repo.Update(med); err != nil {
		return err
	}

	return s.repo.CreateTransaction(
		&models.InventoryTransaction{
			MedicationID: req.MedicationID,
			Type: models.InventoryStockIn,
			Quantity: req.Quantity,
			PreviousStock: previous,
			NewStock: med.QuantityInStock,
			Reference: req.Reference,
			PerformedBy: userID,
		},
	)
}