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

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (u *UserService) Login(request params.LoginUser) *responses.Response {
	user, err := u.userRepository.FindByEmail(request.Email)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: fmt.Sprintf("User %s", err.Error()),
			},
		}
	}

	ok := helpers.ComparePassword([]byte(user.Password), []byte(request.Password))
	if !ok {
		return &responses.Response{
			StatusCode: http.StatusUnauthorized,
			Payload: responses.ErrorResponse{
				Error:   "Unauthorized",
				Message: "User Password not match",
			},
		}
	}

	token := helpers.GenerateToken(user.UserId, user.Email, "user")

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Message: "Login successful",
			Data:    token,
		},
	}
}

func (u *UserService) CreateUser(request params.CreateUser) *responses.Response {
	_, err := u.userRepository.FindByEmail(request.Email)
	if err == nil {
		return &responses.Response{
			StatusCode: http.StatusBadRequest,
			Payload: responses.ErrorResponse{
				Error:   "Bad Request",
				Message: "User Already Registered",
			},
		}
	}

	password := helpers.HashPassword(request.Password)
	user := models.User{
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		City:        request.City,
		Zip:         request.Zip,
		Message:     request.Message,
		Password:    password,
		Username:    request.Username,
		Address:     request.Address,
	}

	userData, err := u.userRepository.Create(&user)
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
			Data: userData,
		},
	}
}

func (u *UserService) UpdateUser(request params.UpdateUser) *responses.Response {
	user, err := u.userRepository.FindById(request.UserId)
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
	if helpers.IsFieldNotNull(request.PhoneNumber) {
		user.PhoneNumber = request.PhoneNumber // update phone number
	}
	if helpers.IsFieldNotNull(request.City) {
		user.City = request.City // update city
	}
	if helpers.IsFieldNotNull(request.Zip) {
		user.Zip = request.Zip // update zip
	}
	if helpers.IsFieldNotNull(request.Message) {
		user.Message = request.Message // update message
	}
	if helpers.IsFieldNotNull(request.Address) {
		user.Address = request.Address // update address
	}

	user, err = u.userRepository.Update(user)
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
			Message: fmt.Sprintf("Success update data with id : %d", user.UserId),
		},
	}
}

func (u *UserService) DeleteUser(userId int) *responses.Response {
	user, err := u.userRepository.FindById(userId)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	user, err = u.userRepository.Delete(user)
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
			Message: fmt.Sprintf("Success delete data with id : %d", user.UserId),
		},
	}
}

func (u *UserService) GetUserById(userId int) *responses.Response {
	user, err := u.userRepository.FindById(userId)
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
			Data: user,
		},
	}
}

func (u *UserService) GetAllUsers() *responses.Response {
	users, err := u.userRepository.FindAllUsers()
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	if len(*users) == 0 {
		return &responses.Response{
			StatusCode: http.StatusOK,
			Payload: responses.SuccessResponse{
				Message: "Data Users is Empty",
			},
		}
	}

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Data: users,
		},
	}
}
