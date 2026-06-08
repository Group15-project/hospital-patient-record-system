package dto

type CreateMedicationRequest struct {
	Name string `json:"name" validate:"required"`

	GenericName string `json:"generic_name"`

	Strength string `json:"strength"`

	Form string `json:"form"`

	UnitPrice float64 `json:"unit_price"`

	QuantityInStock int `json:"quantity_in_stock"`

	ReorderLevel int `json:"reorder_level"`
}

type StockAdjustmentRequest struct {
	MedicationID uint `json:"medication_id" validate:"required"`

	Quantity int `json:"quantity" validate:"required"`

	Reference string `json:"reference"`
}