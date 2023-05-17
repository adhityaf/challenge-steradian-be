package responses

type Order struct {
	OrderId     int    `json:"order_id"`
	PickUpLoc   string `json:"pick_up_loc"`
	DropOffLoc  string `json:"drop_off_loc"`
	PickUpDate  string `json:"pick_up_date"`
	DropOffDate string `json:"drop_off_date"`
	PickUpTime  string `json:"pick_up_time"`
	CarId       int    `json:"car_id"`
	UserId      int    `json:"user_id"`
	AdminId     int    `json:"admin_id"`
}
