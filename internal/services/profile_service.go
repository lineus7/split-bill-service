package services

import (
	"net/http"
	"split-bill-service/internal/repositories"
	"split-bill-service/utils"
)

func NewProfileService(repos *repositories.RepositoriesSet) *ProfileService {
	return &ProfileService{
		userRepository: repos.UserRepository,
	}
}

type ProfileService struct {
	userRepository *repositories.UserRepository
}

func (s *ProfileService) ChangePassword(currentPassword string, newPassword string, email string) error {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return err
	}
	if !utils.CheckPasswordHash(currentPassword, user.Password) {
		return utils.NewAppError(http.StatusBadRequest, "Invalid password", nil)
	}
	user.Password, err = utils.HashPassword(newPassword)
	if err != nil {
		return utils.NewAppError(http.StatusInternalServerError, "Failed to hash password", err)
	}
	_, err = s.userRepository.Update(*user)
	if err != nil {
		return utils.NewAppError(http.StatusInternalServerError, "Failed to update user", err)
	}
	return nil
}
