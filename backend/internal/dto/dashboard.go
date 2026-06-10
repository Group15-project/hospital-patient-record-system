package dto


type DashboardSummary struct {
	TotalPatients      int64   `json:"total_patients"`
	TotalDoctors       int64   `json:"total_doctors"`
	TotalAppointments  int64   `json:"total_appointments"`
	TotalPrescriptions int64   `json:"total_prescriptions"`
	TotalLabRequests   int64   `json:"total_lab_requests"`
	RecentPatients  []PatientDashboardItem `json:"recent_patients"`

	TotalRevenue float64 `json:"total_revenue"`
	AppointmentsToday  int64   `json:"appointments_today"`

	DoctorsAvailable   int64   `json:"doctors_available"`
	EmergencyCases     int64   `json:"emergency_cases"`

	OutstandingBalance float64 `json:"outstanding_balance"`
}

type RevenueReport struct {
	TotalRevenue float64 `json:"total_revenue"`
}

type InventoryAlert struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	QuantityInStock int    `json:"quantity_in_stock"`
	ReorderLevel    int    `json:"reorder_level"`
}
type TodayDashboard struct {
	AppointmentsToday int64 `json:"appointments_today"`

	ConsultationsToday int64 `json:"consultations_today"`

	PaymentsToday int64 `json:"payments_today"`

	RevenueToday float64 `json:"revenue_today"`

	LabRequestsToday int64 `json:"lab_requests_today"`

	PrescriptionsToday int64 `json:"prescriptions_today"`
}

type PatientDashboardItem struct {
	PatientNumber string `json:"patient_number"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	Phone         string `json:"phone"`
	Status        string `json:"status"`
}



