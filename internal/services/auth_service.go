package services

import (
	"net/http"
	"split-bill-service/internal/enums"
	"split-bill-service/internal/models"
	"split-bill-service/internal/repositories"
	"split-bill-service/utils"

	"gorm.io/gorm"
)
func NewAuthService(repos *repositories.RepositoriesSet) *AuthService {
	return &AuthService{
		userRepository: repos.UserRepository,
		geminiRepository: repos.GeminiRepository,
	}
}
type AuthService struct {
	userRepository *repositories.UserRepository
	geminiRepository *repositories.GeminiRepository
}

func (s *AuthService) Login(email string, password string) (*models.User, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewAppError(http.StatusNotFound, "User not found", err)
		}
		return nil, err
	}
	
	if user == nil || !utils.CheckPasswordHash(password, user.Password) {
		return nil, utils.NewAppError(http.StatusUnauthorized, "Invalid Credentials", nil)
	}

	if !user.IsActive {
		return nil, utils.NewAppError(http.StatusUnauthorized, "User Has Not Been Activated", nil)
	}

	user.Password = ""
	return user, nil
}

func (s *AuthService) Register(email string, password string, name string, username string) (*models.User, error) {
	user := models.User{
		Email:    email,
		Password: password,
		Name:     name,
		Username: username,
		RoleID: enums.ROLE_USER,
		IsActive: false,
	}

	data, _ := s.userRepository.GetByEmail(email)
	if data != nil {
		return nil, utils.NewAppError(http.StatusConflict, "Email already exists", nil)
	}

	data, _ = s.userRepository.GetByUsername(username)
	if data != nil {
		return nil, utils.NewAppError(http.StatusConflict, "Username already exists", nil)
	}

	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	newUser.Password = ""
	return newUser, nil
}