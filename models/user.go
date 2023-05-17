package models

type User struct {
	UserId      int    `gorm:"not null;uniqueIndex;primaryKey;" json:"user_id"`
	Email       string `gorm:"not null;uniqueIndex;size:256;" json:"email"`
	PhoneNumber string `gorm:"not null;" json:"phone_number"`
	City        string `gorm:"not null;" json:"city"`
	Zip         string `gorm:"not null;" json:"zip"`
	Message     string `gorm:"not null;" json:"message"`
	Password    string `gorm:"not null;" json:"password"`
	Username    string `gorm:"not null;" json:"username"`
	Address     string `gorm:"not null;" json:"address"`
}
