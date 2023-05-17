package repositories

import (
	"github.com/adhityaf/challenge-steradian-be/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *models.Admin) (*models.Admin, error)
	Update(admin *models.Admin) (*models.Admin, error)
	Delete(admin *models.Admin) (*models.Admin, error)
	FindById(adminId int) (*models.Admin, error)
	FindByEmail(email string) (*models.Admin, error)
	FindAllAdmins() (*[]models.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (a *adminRepository) Create(admin *models.Admin) (*models.Admin, error) {
	err := a.db.Create(&admin).Error
	return admin, err
}

func (a *adminRepository) Update(admin *models.Admin) (*models.Admin, error) {
	err := a.db.Save(&admin).Error
	return admin, err
}

func (a *adminRepository) Delete(admin *models.Admin) (*models.Admin, error) {
	err := a.db.Delete(&admin).Error
	return admin, err
}

func (a *adminRepository) FindById(adminId int) (*models.Admin, error) {
	var admin *models.Admin
	err := a.db.Where("admin_id = ?", adminId).First(&admin).Error
	return admin, err
}

func (a *adminRepository) FindByEmail(email string) (*models.Admin, error) {
	var admin *models.Admin
	err := a.db.Where("email = ?", email).First(&admin).Error
	return admin, err
}

func (a *adminRepository) FindAllAdmins() (*[]models.Admin, error) {
	var admins *[]models.Admin
	err := a.db.Find(&admins).Error
	return admins, err
}
