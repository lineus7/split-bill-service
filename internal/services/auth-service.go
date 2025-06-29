package services

import (
	"errors"
	"split-bill-service/internal/models"
	"split-bill-service/internal/repositories"
	"split-bill-service/utils"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (s *AuthService) Login(email string, password string) (models.User, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return models.User{}, errors.New("invalid credentials")
	}
	user.Password = ""
	return user, nil
}

func (s *AuthService) Register(email string, password string, name string) (models.User, error) {
	user := models.User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	if _, err := s.userRepository.GetByEmail(email); err == nil {
		return models.User{}, errors.New("record already exists")
	}
	user, err := s.userRepository.Create(user)
	user.Password = ""
	return user, err
}

