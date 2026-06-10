package handler

import (
	"fmt"
	"hospital-backend/internal/models"
	"hospital-backend/internal/service"
	"hospital-backend/internal/utils"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedicalDocumentHandler struct {
	service  service.MedicalDocumentService
	validate *validator.Validate
}

func NewMedicalDocumentHandler(
	service service.MedicalDocumentService,
) *MedicalDocumentHandler {
	return &MedicalDocumentHandler{
		service:  service,
		validate: validator.New(),
	}
}


func (h *MedicalDocumentHandler) Upload(
	c *gin.Context,
) {

	file, err := c.FormFile("file")

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"file required",
			nil,
		)
		return
	}

	patientID, _ := strconv.ParseUint(
		c.PostForm("patient_id"),
		10,
		64,
	)

	title := c.PostForm("title")

	documentType := c.PostForm("document_type")

	fileName := fmt.Sprintf(
		"%d_%s",
		time.Now().Unix(),
		file.Filename,
	)

	savePath := filepath.Join(
		"uploads",
		"documents",
		fileName,
	)

	if err := c.SaveUploadedFile(
		file,
		savePath,
	); err != nil {

		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to save file",
			err.Error(),
		)

		return
	}

	userID := c.GetUint("user_id")

	doc := models.MedicalDocument{
		PatientID: uint(patientID),

		Title: title,

		DocumentType: models.MedicalDocumentType(
			documentType,
		),

		FileName: file.Filename,

		FilePath: savePath,

		MimeType: file.Header.Get(
			"Content-Type",
		),

		FileSize: file.Size,

		UploadedBy: userID,
	}

	err = h.service.Upload(
		&doc,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"upload failed",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"document uploaded",
		doc,
	)
}

func (h *MedicalDocumentHandler) GetPatientDocuments(
	c *gin.Context,
) {

	id, _ := strconv.Atoi(
		c.Param("patientId"),
	)

	docs, err := h.service.GetPatientDocuments(
		uint(id),
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to fetch documents",
			err.Error(),
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"documents retrieved",
		docs,
	)
}