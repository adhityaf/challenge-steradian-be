package params

type CreateOrder struct {
	PickUpLoc   string `json:"pick_up_loc" validate:"required"`
	DropOffLoc  string `json:"drop_off_loc" validate:"required"`
	PickUpDate  string `json:"pick_up_date" validate:"required"`
	DropOffDate string `json:"drop_off_date" validate:"required"`
	PickUpTime  string `json:"pick_up_time" validate:"required"`
	CarId       int    `json:"car_id" validate:"required"`
	UserId      int    `json:"user_id" validate:"required"`
	AdminId     int    `json:"admin_id" validate:"required"`
}

type UpdateOrder struct {
	OrderId     int    `json:"order_id" validate:"required"`
	PickUpLoc   string `json:"pick_up_loc"`
	DropOffLoc  string `json:"drop_off_loc"`
	PickUpDate  string `json:"pick_up_date"`
	DropOffDate string `json:"drop_off_date"`
	PickUpTime  string `json:"pick_up_time"`
}
