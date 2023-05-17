package params

type CreateUser struct {
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	City        string `json:"city" validate:"required"`
	Zip         string `json:"zip" validate:"required"`
	Message     string `json:"message" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Address     string `json:"address" validate:"required"`
}

type UpdateUser struct {
	UserId      int    `json:"user_id" validate:"required"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	Message     string `json:"message"`
	Address     string `json:"address"`
}
type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
