package dto

type CreateInvoiceItem struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

type CreateInvoiceRequest struct {
	PatientID uint `json:"patient_id"`

	Items []CreateInvoiceItem `json:"items"`
}

type CreatePaymentRequest struct {
	InvoiceID uint `json:"invoice_id"`

	Amount float64 `json:"amount"`

	Method string `json:"method"`

	Reference string `json:"reference"`
}