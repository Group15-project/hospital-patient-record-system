package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	PaymentCash     PaymentMethod = "CASH"
	PaymentTransfer PaymentMethod = "TRANSFER"
	PaymentPOS      PaymentMethod = "POS"
)

type Payment struct {
	ID uint `gorm:"primaryKey"`

	InvoiceID uint
	Invoice   Invoice

	Amount float64

	Method PaymentMethod `gorm:"type:varchar(20);not null"`

	Reference string

	ReceivedBy uint
	Cashier   User `gorm:"foreignKey:ReceivedBy"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}