package models

type Car struct {
	CarId     int    `gorm:"not null;uniqueIndex;primaryKey;" json:"car_id"`
	Name      string `gorm:"not null;" json:"name"`
	CarType   string `gorm:"not null;" json:"car_type"`
	Rating    int    `gorm:"not null;" json:"rating"`
	Fuel      int    `gorm:"not null;" json:"fuel"`
	Image     string `gorm:"not null;" json:"image"`
	HourRate  int    `gorm:"not null;" json:"hour_rate"`
	DayRate   int    `gorm:"not null;" json:"day_rate"`
	MonthRate int    `gorm:"not null;" json:"month_rate"`
}
