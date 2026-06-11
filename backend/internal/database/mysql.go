package database

import (
	"fmt"
	"log"
	"time"

	"hospital-backend/internal/config"
	//"hospital-backend/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
	return nil, err
}

sqlDB, err := db.DB()
if err != nil {
	return nil, err
}

// Connection pool
sqlDB.SetMaxOpenConns(25)
sqlDB.SetMaxIdleConns(10)
sqlDB.SetConnMaxLifetime(30 * time.Minute)
sqlDB.SetConnMaxIdleTime(10 * time.Minute)

// Ping test
start := time.Now()

err = sqlDB.Ping()
if err != nil {
	return nil, err
}

log.Printf("DB ping took %v", time.Since(start))

return db, nil

	// err = db.AutoMigrate(
	// 	&models.Role{},
	// 	&models.Permission{},
	// 	&models.RolePermission{},

	// 	&models.User{},
	// 	&models.RefreshToken{},
	// 	&models.AuditLog{},

	// 	&models.Patient{},

	// 	&models.Vital{},
	// 	&models.Consultation{},
	// 	&models.Diagnosis{},

	// 	&models.LabRequest{},
	// 	&models.LabResult{},

	// 	&models.Prescription{},
	// 	&models.PrescriptionItem{},
	// 	&models.Appointment{},
	// 	&models.MedicalDocument{},
	// 	&models.Invoice{},
	// 	&models.InvoiceItem{},
	// 	&models.Payment{},
	// 	&models.MedicalRecord{},
	// )
	// if err != nil {
	// 	return nil, err
	// }

	//SeedRolesAndPermissions(db)

	
}
