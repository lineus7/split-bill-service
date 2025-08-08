package services

import (
	"net/http"
	"split-bill-service/internal/models"
	"split-bill-service/internal/repositories"
	"split-bill-service/utils"

	"gorm.io/gorm"
)
func NewAuthService(repos *repositories.RepositoriesSet) *AuthService {
	return &AuthService{userRepository: repos.UserRepository}
}
type AuthService struct {
	userRepository *repositories.UserRepository
}

func (s *AuthService) Login(email string, password string) (*models.User, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewAppError(http.StatusNotFound, "User not found", err)
		}
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, utils.NewAppError(http.StatusUnauthorized, "Invalid Credentials", nil)
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
		return nil, utils.NewAppError(http.StatusConflict, "Email already exists", nil)
	}
	newUser, err := s.userRepository.Create(user)
	newUser.Password = ""
	return newUser, err
}

