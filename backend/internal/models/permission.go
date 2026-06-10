package models

import "gorm.io/gorm"


const (
	// Patients
	PermissionPatientCreate = "patient:create"
	PermissionPatientView   = "patient:view"
	PermissionPatientUpdate = "patient:update"
	PermissionPatientDelete = "patient:delete"

	// Vitals
	PermissionVitalsCreate = "vitals:create"
	PermissionVitalsView   = "vitals:view"
	PermissionVitalsUpdate = "vitals:update"

	// Laboratory
	PermissionLabCreate = "lab:create"
	PermissionLabView   = "lab:view"
	PermissionLabUpdate = "lab:update"

	// Prescription
	PermissionPrescriptionCreate = "prescription:create"
	PermissionPrescriptionView   = "prescription:view"
	PermissionPrescriptionUpdate = "prescription:update"

	// Pharmacy / Inventory
	PermissionInventoryCreate = "inventory:create"
	PermissionInventoryView   = "inventory:view"
	PermissionInventoryUpdate = "inventory:update"
	PermissionInventoryDelete = "inventory:delete"

	// Billing
	PermissionBillingCreate = "billing:create"
	PermissionBillingView   = "billing:view"
	PermissionBillingUpdate = "billing:update"

	// Users
	PermissionUserCreate = "user:create"
	PermissionUserView   = "user:view"
	PermissionUserUpdate = "user:update"
	PermissionUserDelete = "user:delete"

	// Audit Logs
	PermissionAuditView = "audit:view"
)
type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Description string

	Roles []Role `gorm:"many2many:role_permissions"`

	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}