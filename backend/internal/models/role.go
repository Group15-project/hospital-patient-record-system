package models

import "gorm.io/gorm"

type Role struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"unique;not null"`
	Description string

	Users       []User       `gorm:"foreignKey:RoleID"`
	Permissions []Permission `gorm:"many2many:role_permissions"`

	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}