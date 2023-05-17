package repositories

import (
	"github.com/adhityaf/challenge-steradian-be/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) (*models.User, error)
	FindById(userId int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAllUsers() (*[]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userRepository) Update(user *models.User) (*models.User, error) {
	err := u.db.Save(&user).Error
	return user, err
}

func (u *userRepository) Delete(user *models.User) (*models.User, error) {
	err := u.db.Delete(&user).Error
	return user, err
}

func (u *userRepository) FindById(userId int) (*models.User, error) {
	var user *models.User
	err := u.db.Where("user_id = ?", userId).First(&user).Error
	return user, err
}

func (u *userRepository) FindByEmail(email string) (*models.User, error) {
	var user *models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *userRepository) FindAllUsers() (*[]models.User, error) {
	var users *[]models.User
	err := u.db.Find(&users).Error
	return users, err
}
