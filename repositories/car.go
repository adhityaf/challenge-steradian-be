package repositories

import (
	"github.com/adhityaf/challenge-steradian-be/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	Create(car *models.Car) (*models.Car, error)
	Update(car *models.Car) (*models.Car, error)
	Delete(car *models.Car) (*models.Car, error)
	FindById(carId int) (*models.Car, error)
	FindAllCars() (*[]models.Car, error)
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{
		db: db,
	}
}

func (c *carRepository) Create(car *models.Car) (*models.Car, error) {
	err := c.db.Create(&car).Error
	return car, err
}

func (c *carRepository) Update(car *models.Car) (*models.Car, error) {
	err := c.db.Save(&car).Error
	return car, err
}

func (c *carRepository) Delete(car *models.Car) (*models.Car, error) {
	err := c.db.Delete(&car).Error
	return car, err
}

func (c *carRepository) FindById(carId int) (*models.Car, error) {
	var car *models.Car
	err := c.db.Where("car_id = ?", carId).First(&car).Error
	return car, err
}

func (c *carRepository) FindAllCars() (*[]models.Car, error) {
	var cars *[]models.Car
	err := c.db.Find(&cars).Error
	return cars, err
}
