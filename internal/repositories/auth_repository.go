package repositories

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"hnex.com/internal/config"
	"hnex.com/internal/models"
	"hnex.com/internal/utils"
)

type AuthRepository struct {
	DB *gorm.DB
}

func (r *AuthRepository) UpdateRefreshToken(id string, refreshToken *string) error {
	var hashedRefreshToken *string

	if refreshToken == nil {
		config.RedisClient.Del(context.Background(), fmt.Sprintf("user:%s:refresh_token", id))

		hashedRefreshToken = nil
	} else {
		hash, err := utils.HashPassword(*refreshToken)
		if err != nil {
			return err
		}

		config.RedisClient.Set(context.Background(), fmt.Sprintf("user:%s:refresh_token", id), hash, 0)

		hashedRefreshToken = &hash
	}

	return r.DB.Model(&models.User{}).Where("id = ?", id).Update("refresh_token", hashedRefreshToken).Error
}

func (r *AuthRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}
