package repository

import (
	"hospital-backend/internal/dto"
	"hospital-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type DashboardRepository interface {
	GetSummary() (*dto.DashboardSummary, error)
	GetLowStockMedications() (
		[]dto.InventoryAlert,
		error,
	)
	GetTodaySummary() (
		*dto.TodayDashboard,
		error,
	)
}

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(
	db *gorm.DB,
) DashboardRepository {
	return &dashboardRepository{
		db: db,
	}
}

func (r *dashboardRepository) GetSummary() (
	*dto.DashboardSummary,
	error,
) {

	summary := &dto.DashboardSummary{}

	r.db.Model(&models.Patient{}).
		Count(&summary.TotalPatients)

	r.db.Model(&models.User{}).
		Where("role_id IN (?)",
			r.db.
				Model(&models.Role{}).
				Select("id").
				Where("name = ?", models.RoleDoctor),
		).
		Count(&summary.TotalDoctors)

	r.db.Model(&models.Appointment{}).
		Count(&summary.TotalAppointments)
	r.db.Model(&models.Appointment{}).
		Where(
			"priority = ?",
			models.AppointmentEmergency,
		).
		Where(
			"status = ?",
			models.AppointmentScheduled,
		).
		Count(&summary.EmergencyCases)

	r.db.Model(&models.Prescription{}).
		Count(&summary.TotalPrescriptions)

	r.db.Model(&models.LabRequest{}).
		Count(&summary.TotalLabRequests)

	r.db.Model(&models.Payment{}).
		Select("COALESCE(SUM(amount),0)").
		Scan(&summary.TotalRevenue)

	r.db.Model(&models.Invoice{}).
		Select("COALESCE(SUM(balance),0)").
		Scan(&summary.OutstandingBalance)

	var recentPatients []dto.PatientDashboardItem

	err := r.db.
		Model(&models.Patient{}).
		Select(`
		patient_number,
		first_name,
		last_name,
		gender,
		phone,
		is_active
	`).
		Order("created_at DESC").
		Limit(10).
		Scan(&recentPatients).
		Error

	if err != nil {
		return nil, err
	}

	summary.RecentPatients = recentPatients

	return summary, nil
}

func (r *dashboardRepository) GetLowStockMedications() (
	[]dto.InventoryAlert,
	error,
) {

	var alerts []dto.InventoryAlert

	err := r.db.
		Model(&models.Medication{}).
		Select(
			"id, name, quantity_in_stock, reorder_level",
		).
		Where(
			"quantity_in_stock <= reorder_level",
		).
		Find(&alerts).
		Error

	return alerts, err
}

func (r *dashboardRepository) GetTodaySummary() (
	*dto.TodayDashboard,
	error,
) {

	summary := &dto.TodayDashboard{}

	now := time.Now()

	startOfDay := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	endOfDay := startOfDay.Add(24 * time.Hour)

	r.db.Model(&models.Appointment{}).
		Where(
			"created_at >= ? AND created_at < ?",
			startOfDay,
			endOfDay,
		).
		Count(&summary.AppointmentsToday)

	r.db.Model(&models.Consultation{}).
		Where(
			"created_at >= ? AND created_at < ?",
			startOfDay,
			endOfDay,
		).
		Count(&summary.ConsultationsToday)

	r.db.Model(&models.Payment{}).
		Where(
			"created_at >= ? AND created_at < ?",
			startOfDay,
			endOfDay,
		).
		Count(&summary.PaymentsToday)

	r.db.Model(&models.Payment{}).
		Select("COALESCE(SUM(amount),0)").
		Where(
			"created_at >= ? AND created_at < ?",
			startOfDay,
			endOfDay,
		).
		Scan(&summary.RevenueToday)

	r.db.Model(&models.LabRequest{}).
		Where(
			"created_at >= ? AND created_at < ?",
			startOfDay,
			endOfDay,
		).
		Count(&summary.LabRequestsToday)

	r.db.Model(&models.Prescription{}).
		Where(
			"created_at >= ? AND created_at < ?",
			startOfDay,
			endOfDay,
		).
		Count(&summary.PrescriptionsToday)

	return summary, nil
}
