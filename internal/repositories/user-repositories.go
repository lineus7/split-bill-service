package repositories

import (
	"split-bill-service/internal/models"
	"split-bill-service/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil	
}

func (s *UserRepository) Create(user models.User) (models.User, error) {
	user.Password, _ = utils.HashPassword(user.Password)
	if err := s.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserRepository) Update(user models.User) (models.User, error) {
	if err := s.db.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserRepository) Delete(user models.User) (models.User, error) {
	if err := s.db.Delete(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

