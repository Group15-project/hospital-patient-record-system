package repository

import (
	"hospital-backend/internal/models"

	"gorm.io/gorm"
)

type AuditRepository interface {
	Create(*models.AuditLog) error

	List(
		page int,
		limit int,
	) ([]models.AuditLog, error)
}

type auditRepository struct {
	db *gorm.DB
}

func NewAuditRepository(
	db *gorm.DB,
) AuditRepository {
	return &auditRepository{
		db: db,
	}
}

func (r *auditRepository) Create(
	log *models.AuditLog,
) error {

	return r.db.Create(log).Error
}

func (r *auditRepository) List(
	page int,
	limit int,
) ([]models.AuditLog, error) {

	offset := (page - 1) * limit

	var logs []models.AuditLog

	err := r.db.
		Preload("User").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&logs).
		Error

	return logs, err
}
