package services

import (
	"fmt"
	"net/http"

	"github.com/adhityaf/challenge-steradian-be/helpers"
	"github.com/adhityaf/challenge-steradian-be/models"
	"github.com/adhityaf/challenge-steradian-be/params"
	"github.com/adhityaf/challenge-steradian-be/repositories"
	"github.com/adhityaf/challenge-steradian-be/responses"
)

type OrderService struct {
	orderRepository repositories.OrderRepository
	carRepository   repositories.CarRepository
	adminRepository repositories.AdminRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, carRepository repositories.CarRepository, adminRepository repositories.AdminRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		carRepository:   carRepository,
		adminRepository: adminRepository,
	}
}

func (o *OrderService) CreateOrder(request params.CreateOrder) *responses.Response {
	// check car is exist
	_, err := o.carRepository.FindById(request.CarId)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: fmt.Sprintf("Car %s", err.Error()),
			},
		}
	}

	// check admin is exist
	_, err = o.adminRepository.FindById(request.AdminId)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: fmt.Sprintf("Admin %s", err.Error()),
			},
		}
	}

	order := models.Order{
		PickUpLoc:   request.PickUpLoc,
		DropOffLoc:  request.DropOffLoc,
		PickUpDate:  request.PickUpDate,
		DropOffDate: request.DropOffDate,
		PickUpTime:  request.PickUpTime,
		CarId:       request.CarId,
		UserId:      request.UserId,
		AdminId:     request.AdminId,
	}

	orderData, err := o.orderRepository.Create(&order)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		}
	}

	return &responses.Response{
		StatusCode: http.StatusCreated,
		Payload: responses.SuccessResponse{
			Data: responses.Order{
				OrderId:     orderData.OrderId,
				PickUpLoc:   orderData.PickUpLoc,
				DropOffLoc:  orderData.DropOffLoc,
				PickUpDate:  orderData.PickUpDate,
				DropOffDate: orderData.DropOffDate,
				PickUpTime:  orderData.PickUpTime,
				CarId:       orderData.CarId,
				UserId:      orderData.UserId,
				AdminId:     orderData.AdminId,
			},
		},
	}
}

func (o *OrderService) UpdateOrder(request params.UpdateOrder) *responses.Response {
	order, err := o.orderRepository.FindById(request.OrderId)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	// IF field is not null then update
	if helpers.IsFieldNotNull(request.PickUpLoc) {
		order.PickUpLoc = request.PickUpLoc // update pick up location
	}
	if helpers.IsFieldNotNull(request.DropOffLoc) {
		order.DropOffLoc = request.DropOffLoc // update drop off location
	}
	if helpers.IsFieldNotNull(request.PickUpDate) {
		order.PickUpDate = request.PickUpDate // update pick up date
	}
	if helpers.IsFieldNotNull(request.DropOffDate) {
		order.DropOffDate = request.DropOffDate // update drop off date
	}
	if helpers.IsFieldNotNull(request.PickUpTime) {
		order.PickUpTime = request.PickUpTime // update pick up time
	}

	order, err = o.orderRepository.Update(order)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		}
	}

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Message: fmt.Sprintf("Success update data with id : %d", order.OrderId),
		},
	}
}

func (o *OrderService) DeleteOrder(orderId int) *responses.Response {
	order, err := o.orderRepository.FindById(orderId)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	order, err = o.orderRepository.Delete(order)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: err.Error(),
			},
		}
	}

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Message: fmt.Sprintf("Success delete data with id : %d", order.OrderId),
		},
	}
}

func (o *OrderService) GetOrderById(orderId int) *responses.Response {
	order, err := o.orderRepository.FindById(orderId)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Data: responses.Order{
				OrderId:     order.OrderId,
				PickUpLoc:   order.PickUpLoc,
				DropOffLoc:  order.DropOffLoc,
				PickUpDate:  order.PickUpDate,
				DropOffDate: order.DropOffDate,
				PickUpTime:  order.PickUpTime,
				CarId:       order.CarId,
				UserId:      order.UserId,
				AdminId:     order.AdminId,
			},
		},
	}
}

func (o *OrderService) GetAllOrders() *responses.Response {
	orderData, err := o.orderRepository.FindAllOrders()
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	if len(*orderData) == 0 {
		return &responses.Response{
			StatusCode: http.StatusOK,
			Payload: responses.SuccessResponse{
				Message: "Data Orders is Empty",
			},
		}
	}

	var orders []responses.Order
	for _, order := range *orderData {
		orders = append(orders, responses.Order{
			OrderId:     order.OrderId,
			PickUpLoc:   order.PickUpLoc,
			DropOffLoc:  order.DropOffLoc,
			PickUpDate:  order.PickUpDate,
			DropOffDate: order.DropOffDate,
			PickUpTime:  order.PickUpTime,
			CarId:       order.CarId,
			UserId:      order.UserId,
			AdminId:     order.AdminId,
		})
	}

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Data: orders,
		},
	}
}
