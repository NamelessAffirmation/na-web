package services

import (
	"github.com/snowlynxsoftware/na-api/server/database/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
	TokenService   *TokenService
}

func NewUserService(userRepository *repositories.UserRepository, tokenService *TokenService) *UserService {
	return &UserService{UserRepository: userRepository, TokenService: tokenService}
}

func (s *UserService) GetUserById(id int) (*repositories.UserEntity, error) {
	return s.UserRepository.GetUserById(id)
}
