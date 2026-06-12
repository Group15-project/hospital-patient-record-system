package models

import (
	"time"

	"gorm.io/gorm"
)

type LabResult struct {
	ID uint `gorm:"primaryKey"`

	LabRequestID uint       `gorm:"not null;uniqueIndex"`
	LabRequest   LabRequest `gorm:"foreignKey:LabRequestID"`

	Result string `gorm:"type:text"`

	Remarks string `gorm:"type:text"`

	UploadedBy uint
	LabTech    User `gorm:"foreignKey:UploadedBy"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}