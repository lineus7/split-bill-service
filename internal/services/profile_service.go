package services

import (
	"net/http"
	"split-bill-service/internal/models"
	"split-bill-service/internal/repositories"
	"split-bill-service/utils"

	"gorm.io/gorm"
)
func NewProfileService(repos *repositories.RepositoriesSet) *ProfileService {
	return &ProfileService{
		userRepository: repos.UserRepository,
	}
}
type ProfileService struct {
	userRepository *repositories.UserRepository
}

func (s *ProfileService) ChangePassword(email string, oldPassword string, newPassword string) (*models.User, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewAppError(http.StatusNotFound, "User not found", err)
		}
		return nil, err
	}

	if user == nil || !utils.CheckPasswordHash(oldPassword, user.Password) {
		return nil, utils.NewAppError(http.StatusUnauthorized, "Invalid Old Password", nil)
	}

	user.Password, err = utils.HashPassword(newPassword)
	if err != nil {
		return nil, err
	}
	
	_, err = s.userRepository.Update(*user)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}
