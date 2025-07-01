package services

import (
	"net/http"
	"split-bill-service/internal/models"
	"split-bill-service/internal/repositories"
	"split-bill-service/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (s *AuthService) Login(email string, password string) (*models.User, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appErr := utils.AppError{
				Code:    http.StatusNotFound,
				Message: "Record Not Found",
				Err:     err,
			}
			return nil, &appErr
		}
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		appErr := utils.AppError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Credentials",
		}
		return nil, &appErr
	}
	user.Password = ""
	return user, nil
}

func (s *AuthService) Register(email string, password string, name string) (*models.User, error) {
	user := models.User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	if _, err := s.userRepository.GetByEmail(email); err == nil {
		appErr := utils.AppError{
			Code:    http.StatusConflict,
			Message: "Email already exists",
			Err:     err,
		}
		return nil, &appErr
	}
	newUser, err := s.userRepository.Create(user)
	newUser.Password = ""
	return newUser, err
}

