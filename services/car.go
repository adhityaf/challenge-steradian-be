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

type CarService struct {
	carRepository repositories.CarRepository
}

func NewCarService(repo repositories.CarRepository) *CarService {
	return &CarService{
		carRepository: repo,
	}
}

func (c *CarService) CreateCar(request params.CreateCar) *responses.Response {
	car := models.Car{
		Name:      request.Name,
		CarType:   request.CarType,
		Rating:    request.Rating,
		Fuel:      request.Fuel,
		Image:     request.Image,
		HourRate:  request.HourRate,
		DayRate:   request.DayRate,
		MonthRate: request.MonthRate,
	}

	carData, err := c.carRepository.Create(&car)
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
			Data: carData,
		},
	}
}

func (c *CarService) UpdateCar(request params.UpdateCar) *responses.Response {
	car, err := c.carRepository.FindById(request.CarId)
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
	if helpers.IsFieldNotNull(request.Name) {
		car.Name = request.Name // update name
	}
	if helpers.IsFieldNotNull(request.CarType) {
		car.CarType = request.CarType // update car type
	}
	if helpers.IsFieldNotNull(request.Rating) {
		car.Rating = request.Rating // update rating
	}
	if helpers.IsFieldNotNull(request.Fuel) {
		car.Fuel = request.Fuel // update fuel
	}
	if helpers.IsFieldNotNull(request.Image) {
		car.Image = request.Image // update image
	}
	if helpers.IsFieldNotNull(request.HourRate) {
		car.HourRate = request.HourRate // update hour rate
	}
	if helpers.IsFieldNotNull(request.DayRate) {
		car.DayRate = request.DayRate // update day rate
	}
	if helpers.IsFieldNotNull(request.MonthRate) {
		car.MonthRate = request.MonthRate // update month rate
	}

	car, err = c.carRepository.Update(car)
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
			Message: fmt.Sprintf("Success update data with id : %d", car.CarId),
		},
	}
}

func (c *CarService) DeleteCar(carId int) *responses.Response {
	car, err := c.carRepository.FindById(carId)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	car, err = c.carRepository.Delete(car)
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
			Message: fmt.Sprintf("Success delete data with id : %d", car.CarId),
		},
	}
}

func (c *CarService) GetCarById(carId int) *responses.Response {
	car, err := c.carRepository.FindById(carId)
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
			Data: car,
		},
	}
}

func (c *CarService) GetAllCars() *responses.Response {
	cars, err := c.carRepository.FindAllCars()
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	if len(*cars) == 0 {
		return &responses.Response{
			StatusCode: http.StatusOK,
			Payload: responses.SuccessResponse{
				Message: "Data Cars is Empty",
			},
		}
	}

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Data: cars,
		},
	}
}
