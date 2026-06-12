package handler

import (
	"net/http"
	"strconv"

	"hospital-backend/internal/dto"
	"hospital-backend/internal/service"
	"hospital-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService service.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(
	authService service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validator.New(),
	}
}

// Login godoc
//
//	@Summary		User Login
//	@Description	Authenticate user and return JWT access token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Login credentials"
//	@Success		200		{object}	service.LoginResponse
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		401		{object}	map[string]interface{}
//	@Router			/api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"validation failed",
			err.Error(),
		)
		return
	}

	auditCtx := utils.NewAuditContext(c)

	response, err := h.authService.Login(
		req.Email,
		req.Password,
		auditCtx,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusUnauthorized,
			err.Error(),
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		http.StatusOK,
		"login successful",
		response,
	)
}

func (h *AuthHandler) RegisterUser(
	c *gin.Context,
) {

	var req dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid request body",
			err.Error(),
		)

		return
	}

	if err := h.validate.Struct(req); err != nil {

		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"validation failed",
			err.Error(),
		)

		return
	}

	auditCtx := utils.NewAuditContext(c)

	err := h.authService.RegisterUser(
		req,
		auditCtx,
	)

	if err != nil {

		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)

		return
	}

	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"user created successfully",
		nil,
	)
}
func (h *AuthHandler) GetDoctors(
	c *gin.Context,
) {

	doctors, err :=
		h.authService.GetDoctors()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data": doctors,
		},
	)
}

func (h *AuthHandler) Profile(
	c *gin.Context,
) {

	userID := c.GetUint("user_id")

	user, err := h.authService.GetProfile(
		userID,
	)

	if err != nil {

		utils.ErrorResponse(
			c,
			http.StatusNotFound,
			"user not found",
			nil,
		)

		return
	}

	profile := dto.ProfileResponse{
	ID:        user.ID,
	FirstName: user.FirstName,
	LastName:  user.LastName,
	Email:     user.Email,
	Phone:     user.Phone,
	Role:      user.Role.Name,
}
utils.SuccessResponse(
	c,
	http.StatusOK,
	"profile retrieved successfully",
	profile,
)
}

func (h *AuthHandler) GetStaff(
    c *gin.Context,
) {

    users, err :=
        h.authService.GetStaff()

    if err != nil {

        c.JSON(
            http.StatusInternalServerError,
            gin.H{
                "success": false,
                "message": err.Error(),
            },
        )

        return
    }

    c.JSON(
        http.StatusOK,
        gin.H{
            "success": true,
            "data": users,
        },
    )
}

func (h *AuthHandler) DeleteUser(
    c *gin.Context,
) {

    id, err :=
        strconv.ParseUint(
            c.Param("id"),
            10,
            64,
        )

    if err != nil {

        c.JSON(
            http.StatusBadRequest,
            gin.H{
                "success": false,
                "message": "invalid id",
            },
        )

        return
    }

    err =
        h.authService.DeleteUser(
            uint(id),
        )

    if err != nil {

        c.JSON(
            http.StatusInternalServerError,
            gin.H{
                "success": false,
                "message": err.Error(),
            },
        )

        return
    }

    c.JSON(
        http.StatusOK,
        gin.H{
            "success": true,
            "message": "user deleted",
        },
    )
}

func (h *AuthHandler) GetRoles(
	c *gin.Context,
) {

	roles, err :=
		h.authService.GetRoles()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"data": roles,
		},
	)
}