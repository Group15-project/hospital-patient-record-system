package service

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/repository"
)



type DashboardService interface {

	GetSummary() (
		*dto.DashboardSummary,
		error,
	)
	GetLowStockMedications() (
		[]dto.InventoryAlert,
		error,
	)
	GetTodaySummary() (
	*dto.TodayDashboard,
	error,
)

}

type dashboardService struct {
	repo repository.DashboardRepository
}

func NewDashboardService(
	repo repository.DashboardRepository,
) DashboardService {
	return &dashboardService{
		repo: repo,
	}
}

func (s *dashboardService) GetSummary() (
	*dto.DashboardSummary,
	error,
) {
	return s.repo.GetSummary()
}

func (s *dashboardService) GetLowStockMedications() (
	[]dto.InventoryAlert,
	error,
) {
	return s.repo.GetLowStockMedications()
}

func (s *dashboardService) GetTodaySummary() (
	*dto.TodayDashboard,
	error,
) {
	return s.repo.GetTodaySummary()
}