package params

type CreateCar struct {
	Name      string `json:"name" validate:"required"`
	CarType   string `json:"car_type" validate:"required"`
	Rating    int    `json:"rating" validate:"required"`
	Fuel      int    `json:"fuel" validate:"required"`
	Image     string `json:"image" validate:"required"`
	HourRate  int    `json:"hour_rate" validate:"required"`
	DayRate   int    `json:"day_rate" validate:"required"`
	MonthRate int    `json:"month_rate" validate:"required"`
}

type UpdateCar struct {
	CarId     int    `json:"car_id" validate:"required"`
	Name      string `json:"name"`
	CarType   string `json:"car_type"`
	Rating    int    `json:"rating"`
	Fuel      int    `json:"fuel"`
	Image     string `json:"image"`
	HourRate  int    `json:"hour_rate"`
	DayRate   int    `json:"day_rate"`
	MonthRate int    `json:"month_rate"`
}
