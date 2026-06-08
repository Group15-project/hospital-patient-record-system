package database

import (
	"log"

	"hospital-backend/internal/models"
	"hospital-backend/internal/utils"

	"gorm.io/gorm"
)

func SeedRolesAndPermissions(db *gorm.DB) {
	permissions := []string{
		models.PermissionPatientCreate,
		models.PermissionPatientView,
		models.PermissionPatientUpdate,
		models.PermissionPatientDelete,

		models.PermissionVitalsCreate,
		models.PermissionVitalsView,
		models.PermissionVitalsUpdate,

		models.PermissionLabCreate,
		models.PermissionLabView,
		models.PermissionLabUpdate,

		models.PermissionPrescriptionCreate,
		models.PermissionPrescriptionView,
		models.PermissionPrescriptionUpdate,

		models.PermissionInventoryCreate,
		models.PermissionInventoryView,
		models.PermissionInventoryUpdate,
		models.PermissionInventoryDelete,

		models.PermissionBillingCreate,
		models.PermissionBillingView,
		models.PermissionBillingUpdate,

		models.PermissionUserCreate,
		models.PermissionUserView,
		models.PermissionUserUpdate,
		models.PermissionUserDelete,

		models.PermissionAuditView,
	}

	for _, permission := range permissions {
		db.FirstOrCreate(
			&models.Permission{},
			models.Permission{
				Name: permission,
			},
		)
	}

	seedRoles(db)
seedSuperAdmin(db)

log.Println("roles, permissions and admin seeded")

	log.Println("roles and permissions seeded")
}

func seedRoles(db *gorm.DB) {
	rolePermissions := map[string][]string{

		models.RoleSuperAdmin: {},

		models.RoleAdmin: {
			models.PermissionUserCreate,
			models.PermissionUserView,
			models.PermissionUserUpdate,
			models.PermissionUserDelete,

			models.PermissionPatientView,

			models.PermissionAuditView,
		},

		models.RoleReceptionist: {
			models.PermissionPatientCreate,
			models.PermissionPatientView,
			models.PermissionPatientUpdate,

			models.PermissionBillingCreate,
			models.PermissionBillingView,
		},

		models.RoleNurse: {
			models.PermissionPatientView,

			models.PermissionVitalsCreate,
			models.PermissionVitalsView,
			models.PermissionVitalsUpdate,
		},

		models.RoleDoctor: {
			models.PermissionPatientView,

			models.PermissionVitalsView,

			models.PermissionLabView,

			models.PermissionPrescriptionCreate,
			models.PermissionPrescriptionView,
			models.PermissionPrescriptionUpdate,
		},

		models.RoleLabTechnician: {
			models.PermissionPatientView,

			models.PermissionLabCreate,
			models.PermissionLabView,
			models.PermissionLabUpdate,
		},

		models.RolePharmacist: {
			models.PermissionPrescriptionView,

			models.PermissionInventoryView,
			models.PermissionInventoryUpdate,

			models.PermissionBillingView,
		},

		models.RoleBilling: {
			models.PermissionBillingCreate,
			models.PermissionBillingView,
			models.PermissionBillingUpdate,
		},
	}

	for roleName, permissionNames := range rolePermissions {

		role := models.Role{
			Name: roleName,
		}

		db.FirstOrCreate(&role, role)

		if len(permissionNames) == 0 {
			continue
		}

		var permissions []models.Permission

		db.Where("name IN ?", permissionNames).
			Find(&permissions)

		db.Model(&role).
			Association("Permissions").
			Replace(&permissions)
	}
}
func seedSuperAdmin(db *gorm.DB) {

	var existing models.User

	if err := db.
		Where("email = ?", "admin@hospital.com").
		First(&existing).
		Error; err == nil {
		return
	}

	var superAdminRole models.Role

	if err := db.
		Where("name = ?", models.RoleSuperAdmin).
		First(&superAdminRole).
		Error; err != nil {
		log.Printf("failed to find super admin role: %v", err)
		return
	}

	password, err := utils.HashPassword("Admin@123")
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return
	}

	admin := models.User{
		FirstName:      "Super",
		LastName:       "Admin",
		Email:          "admin@hospital.com",
		HashedPassword: password,
		IsActive:       true,
		RoleID:         superAdminRole.ID,
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Printf("failed to create super admin: %v", err)
		return
	}

	log.Println("super admin seeded")
}