package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adhityaf/challenge-steradian-be/params"
	"github.com/adhityaf/challenge-steradian-be/responses"
	"github.com/adhityaf/challenge-steradian-be/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{
		adminService: *adminService,
	}
}

func (a *AdminController) Login(ctx *gin.Context) {
	var req params.LoginAdmin

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	var validationMessages []string
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage := fmt.Sprintf("Field %s %s.", err.Field(), err.Tag())
			validationMessages = append(validationMessages, validationMessage)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: validationMessages,
			},
		})

		return
	}

	result := a.adminService.Login(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (a *AdminController) CreateAdmin(ctx *gin.Context) {
	var req params.CreateAdmin

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	var validationMessages []string
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage := fmt.Sprintf("Field %s %s.", err.Field(), err.Tag())
			validationMessages = append(validationMessages, validationMessage)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: validationMessages,
			},
		})

		return
	}

	result := a.adminService.CreateAdmin(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (a *AdminController) GetAllAdmins(ctx *gin.Context) {
	result := a.adminService.GetAllAdmins()

	ctx.JSON(result.StatusCode, result.Payload)
}

func (a *AdminController) GetAdminById(ctx *gin.Context) {
	adminId, err := strconv.Atoi(ctx.Param("adminId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	result := a.adminService.GetAdminById(adminId)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (a *AdminController) UpdatePasswordAdmin(ctx *gin.Context) {
	var req params.UpdatePasswordAdmin

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	// Get admin_id from param
	adminId, err := strconv.Atoi(ctx.GetString("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	req.AdminId = adminId

	var validationMessages []string
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage := fmt.Sprintf("Field %s %s.", err.Field(), err.Tag())
			validationMessages = append(validationMessages, validationMessage)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: validationMessages,
			},
		})

		return
	}

	result := a.adminService.UpdatePasswordAdmin(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (a *AdminController) DeleteAdminById(ctx *gin.Context) {
	// Get admin_id from param
	adminId, err := strconv.Atoi(ctx.Param("adminId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	result := a.adminService.DeleteAdmin(adminId)

	ctx.JSON(result.StatusCode, result.Payload)
}
