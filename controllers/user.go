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

type UserController struct {
	userService services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: *userService,
	}
}

func (u *UserController) Login(ctx *gin.Context) {
	var req params.LoginUser

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

	result := u.userService.Login(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (u *UserController) Register(ctx *gin.Context) {
	var req params.CreateUser

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

	result := u.userService.CreateUser(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (u *UserController) GetUserProfile(ctx *gin.Context) {
	// Get user_id from token
	userId, err := strconv.Atoi(ctx.GetString("id"))
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

	result := u.userService.GetUserById(userId)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (u *UserController) GetAllUsers(ctx *gin.Context) {
	result := u.userService.GetAllUsers()

	ctx.JSON(result.StatusCode, result.Payload)
}

func (u *UserController) UpdateUserProfile(ctx *gin.Context) {
	var req params.UpdateUser

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

	// Get user_id from token
	userId, err := strconv.Atoi(ctx.GetString("id"))
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

	// Request from json only name and role
	req.UserId = userId

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

	result := u.userService.UpdateUser(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (u *UserController) DeleteUserById(ctx *gin.Context) {
	// Get user_id from param
	userId, err := strconv.Atoi(ctx.Param("userId"))
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

	result := u.userService.DeleteUser(userId)

	ctx.JSON(result.StatusCode, result.Payload)
}
