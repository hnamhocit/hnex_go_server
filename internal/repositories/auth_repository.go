package repositories

import (
	"gorm.io/gorm"
	"hnex.com/internal/models"
)

type AuthRepository struct {
	DB *gorm.DB
}

func (r *AuthRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
