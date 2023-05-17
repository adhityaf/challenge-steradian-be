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

type CarController struct {
	carService services.CarService
}

func NewCarController(carService *services.CarService) *CarController {
	return &CarController{
		carService: *carService,
	}
}

func (c *CarController) CreateCar(ctx *gin.Context) {
	var req params.CreateCar

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode:  http.StatusBadRequest,
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
			StatusCode:  http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: validationMessages,
			},
		})

		return
	}

	result := c.carService.CreateCar(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (c *CarController) GetAllCars(ctx *gin.Context) {
	result := c.carService.GetAllCars()

	ctx.JSON(result.StatusCode, result.Payload)
}

func (c *CarController) GetCarById(ctx *gin.Context) {
	carId, err := strconv.Atoi(ctx.Param("carId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode:  http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	result := c.carService.GetCarById(carId)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (c *CarController) UpdateCarById(ctx *gin.Context) {
	var req params.UpdateCar

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode:  http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	// Get car_id from param
	carId, err := strconv.Atoi(ctx.Param("carId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode:  http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	// Request from json only name and role
	req.CarId = carId

	var validationMessages []string
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage := fmt.Sprintf("Field %s %s.", err.Field(), err.Tag())
			validationMessages = append(validationMessages, validationMessage)
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode:  http.StatusBadRequest,
			Payload: validationMessages,
		})

		return
	}

	result := c.carService.UpdateCar(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (c *CarController) DeleteCarById(ctx *gin.Context) {
	// Get car_id from param
	carId, err := strconv.Atoi(ctx.Param("carId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			StatusCode:  http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		})

		return
	}

	result := c.carService.DeleteCar(carId)

	ctx.JSON(result.StatusCode, result.Payload)
}
