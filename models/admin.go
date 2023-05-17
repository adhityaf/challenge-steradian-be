package models

type Admin struct {
	AdminId  int    `gorm:"not null;uniqueIndex;primaryKey;" json:"admin_id"`
	Email    string `gorm:"not null;uniqueIndex;size:256;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
}
