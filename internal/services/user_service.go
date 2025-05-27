package services

import "hnex.com/internal/repositories"

type UserService struct {
	Repository repositories.UserRepository
}
