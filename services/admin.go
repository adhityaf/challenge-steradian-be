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

type AdminService struct {
	adminRepository repositories.AdminRepository
}

func NewAdminService(repo repositories.AdminRepository) *AdminService {
	return &AdminService{
		adminRepository: repo,
	}
}

func (a *AdminService) Login(request params.LoginAdmin) *responses.Response {
	admin, err := a.adminRepository.FindByEmail(request.Email)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: fmt.Sprintf("Admin %s", err.Error()),
			},
		}
	}

	ok := helpers.ComparePassword([]byte(admin.Password), []byte(request.Password))
	if !ok {
		return &responses.Response{
			StatusCode: http.StatusUnauthorized,
			Payload: responses.ErrorResponse{
				Error:   "Unauthorized",
				Message: "Password not match",
			},
		}
	}

	token := helpers.GenerateToken(admin.AdminId, admin.Email, "admin")

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Message: "Login successful",
			Data:    token,
		},
	}
}

func (a *AdminService) CreateAdmin(request params.CreateAdmin) *responses.Response {
	password := helpers.HashPassword(request.Password)
	admin := models.Admin{
		Email:    request.Email,
		Password: password,
	}

	adminData, err := a.adminRepository.Create(&admin)
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
			Data: adminData,
		},
	}
}

func (a *AdminService) UpdatePasswordAdmin(request params.UpdatePasswordAdmin) *responses.Response {
	admin, err := a.adminRepository.FindById(request.AdminId)
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
	if helpers.IsFieldNotNull(request.Password) {
		password := helpers.HashPassword(request.Password)
		admin.Password = password // update password
	}

	admin, err = a.adminRepository.Update(admin)
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
			Message: "Success update password Admin",
		},
	}
}

func (a *AdminService) DeleteAdmin(adminId int) *responses.Response {
	admin, err := a.adminRepository.FindById(adminId)
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	admin, err = a.adminRepository.Delete(admin)
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
			Message: fmt.Sprintf("Success delete data with id : %d", admin.AdminId),
		},
	}
}

func (a *AdminService) GetAdminById(adminId int) *responses.Response {
	admin, err := a.adminRepository.FindById(adminId)
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
			Data: admin,
		},
	}
}

func (a *AdminService) GetAllAdmins() *responses.Response {
	admins, err := a.adminRepository.FindAllAdmins()
	if err != nil {
		return &responses.Response{
			StatusCode: http.StatusNotFound,
			Payload: responses.ErrorResponse{
				Error:   "Not Found",
				Message: err.Error(),
			},
		}
	}

	if len(*admins) == 0 {
		return &responses.Response{
			StatusCode: http.StatusOK,
			Payload: responses.SuccessResponse{
				Message: "Data Admins is Empty",
			},
		}
	}

	return &responses.Response{
		StatusCode: http.StatusOK,
		Payload: responses.SuccessResponse{
			Data: admins,
		},
	}
}
