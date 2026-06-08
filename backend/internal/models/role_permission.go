package models

import "github.com/google/uuid"

type RolePermission struct {
	RoleID uuid.UUID `gorm:"primaryKey"`
	PermissionID uuid.UUID `gorm:"primaryKey"`
}