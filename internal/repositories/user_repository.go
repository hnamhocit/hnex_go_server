package repositories

import (
	"gorm.io/gorm"
	"hnex.com/internal/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) UpdateById(id uint32, user *models.User) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *UserRepository) FindById(id uint32) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
