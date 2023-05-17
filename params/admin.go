package params

type CreateAdmin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatePasswordAdmin struct {
	AdminId  int    `json:"admin_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginAdmin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
