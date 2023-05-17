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

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{
		orderService: *orderService,
	}
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var req params.CreateOrder

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
			Payload:    validationMessages,
		})

		return
	}

	result := o.orderService.CreateOrder(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (o *OrderController) GetAllOrders(ctx *gin.Context) {
	result := o.orderService.GetAllOrders()

	ctx.JSON(result.StatusCode, result.Payload)
}

func (o *OrderController) GetOrderById(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
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

	result := o.orderService.GetOrderById(orderId)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (o *OrderController) UpdateOrderById(ctx *gin.Context) {
	var req params.UpdateOrder
	validate := validator.New()

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

	// Get order_id from param
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
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
	req.OrderId = orderId

	var validationMessages []string

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

	result := o.orderService.UpdateOrder(req)

	ctx.JSON(result.StatusCode, result.Payload)
}

func (o *OrderController) DeleteOrderById(ctx *gin.Context) {
	// Get order_id from param
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
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

	result := o.orderService.DeleteOrder(orderId)

	ctx.JSON(result.StatusCode, result.Payload)
}
