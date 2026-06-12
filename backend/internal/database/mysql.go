package database

import (
	"fmt"
	"log"
	"time"

	"hospital-backend/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {

	var dsn string

	if cfg.DBPassword == "" {
		dsn = fmt.Sprintf(
			"postgres://%s@%s:%s/%s?sslmode=require",
			cfg.DBUser,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		)
	} else {
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=require",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		)
	}

	log.Printf("DSN=%s", dsn)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
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

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Printf("DB ping took %v", time.Since(start))

	return db, nil


	// err = db.AutoMigrate(
	// 	&models.Role{},
	// 	&models.Permission{},
	// 	&models.RolePermission{},
	//
	// 	&models.User{},
	// 	&models.RefreshToken{},
	// 	&models.AuditLog{},
	//
	// 	&models.Patient{},
	//
	// 	&models.Vital{},
	// 	&models.Consultation{},
	// 	&models.Diagnosis{},
	//
	// 	&models.LabRequest{},
	// 	&models.LabResult{},
	//
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
	//
	// SeedRolesAndPermissions(db)
}
