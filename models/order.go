package models

type Order struct {
	OrderId     int    `gorm:"not null;uniqueIndex;primaryKey;" json:"order_id"`
	PickUpLoc   string `gorm:"not null;" json:"pick_up_loc"`
	DropOffLoc  string `gorm:"not null;" json:"drop_off_loc"`
	PickUpDate  string `gorm:"not null;" json:"pick_up_date"`
	DropOffDate string `gorm:"not null;" json:"drop_off_date"`
	PickUpTime  string `gorm:"not null;" json:"pick_up_time"`

	// Belongs To
	CarId int `gorm:"not null;" json:"car_id"`
	Car   Car `gorm:"references:CarId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserId int  `gorm:"not null;" json:"user_id"`
	User   User `gorm:"references:CarId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	AdminId int   `gorm:"not null;" json:"admin_id"`
	Admin   Admin `gorm:"references:AdminId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
