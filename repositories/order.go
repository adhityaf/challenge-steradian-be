package repositories

import (
	"github.com/adhityaf/challenge-steradian-be/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) (*models.Order, error)
	Update(order *models.Order) (*models.Order, error)
	Delete(order *models.Order) (*models.Order, error)
	FindById(orderId int) (*models.Order, error)
	FindAllOrders() (*[]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (o *orderRepository) Create(order *models.Order) (*models.Order, error) {
	err := o.db.Create(&order).Error
	return order, err
}

func (o *orderRepository) Update(order *models.Order) (*models.Order, error) {
	err := o.db.Save(&order).Error
	return order, err
}

func (o *orderRepository) Delete(order *models.Order) (*models.Order, error) {
	err := o.db.Delete(&order).Error
	return order, err
}

func (o *orderRepository) FindById(orderId int) (*models.Order, error) {
	var order *models.Order
	err := o.db.Where("order_id = ?", orderId).First(&order).Error
	return order, err
}

func (o *orderRepository) FindAllOrders() (*[]models.Order, error) {
	var orders *[]models.Order
	err := o.db.Find(&orders).Error
	return orders, err
}
