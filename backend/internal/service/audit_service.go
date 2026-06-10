package service

import (
	"hospital-backend/internal/models"
	"hospital-backend/internal/repository"
	"hospital-backend/internal/utils"
)

type AuditService interface {
	Log(
		ctx utils.AuditContext,
		action string,
		resource string,
		resourceID string,
		details string,
	)

	List(
		page int,
		limit int,
	) ([]models.AuditLog, error)
}

type auditService struct {
	repo repository.AuditRepository
}

func NewAuditService(
	repo repository.AuditRepository,
) AuditService {

	return &auditService{
		repo: repo,
	}
}

func (s *auditService) Log(
	ctx utils.AuditContext,
	action string,
	resource string,
	resourceID string,
	details string,
) {

	log := models.AuditLog{
		UserID: ctx.UserID,

		Action: action,

		Resource: resource,

		ResourceID: resourceID,

		Details: details,

		IPAddress: ctx.IPAddress,

		UserAgent: ctx.UserAgent,
	}

	_ = s.repo.Create(&log)
}

func (s *auditService) List(
	page int,
	limit int,
) ([]models.AuditLog, error) {

	return s.repo.List(
		page,
		limit,
	)
}
