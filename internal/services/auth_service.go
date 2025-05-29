package services

import (
	"context"
	"fmt"

	"hnex.com/internal/config"
	"hnex.com/internal/models"
	"hnex.com/internal/repositories"
	"hnex.com/internal/utils"
)

type AuthService struct {
	Repository     repositories.AuthRepository
	UserRepository repositories.UserRepository
}

func (s *AuthService) UpdateRefreshToken(id uint32, refreshToken *string) error {
	var hashedRefreshToken *string

	if refreshToken == nil {
		config.RedisClient.Del(context.Background(), fmt.Sprintf("user:%d:refresh_token", id))

		hashedRefreshToken = nil
	} else {
		hash, err := utils.HashPassword(*refreshToken)
		if err != nil {
			return err
		}

		config.RedisClient.Set(context.Background(), fmt.Sprintf("user:%d:refresh_token", id), hash, 0)

		hashedRefreshToken = &hash
	}

	return s.UserRepository.UpdateById(id, &models.User{
		RefreshToken: hashedRefreshToken,
	})
}
