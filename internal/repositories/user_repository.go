package repositories

import (
	"split-bill-service/config"
	"split-bill-service/internal/models"
	"split-bill-service/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(connection *config.Connection) *UserRepository {
	return &UserRepository{db: connection.DB}
}

func (s *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).Preload("Role").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil	
}

func (s *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserRepository) Create(user models.User) (*models.User, error) {
	user.Password, _ = utils.HashPassword(user.Password)
	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserRepository) Update(user models.User) (*models.User, error) {
	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserRepository) Delete(user models.User) (*models.User, error) {
	if err := s.db.Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

