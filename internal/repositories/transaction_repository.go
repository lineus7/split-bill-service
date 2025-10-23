package repositories

import (
	"split-bill-service/config"
	"split-bill-service/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(connection *config.Connection) *TransactionRepository {
	return &TransactionRepository{db: connection.DB}
}

func (s *TransactionRepository) GetListByUserId(userId uint, search string) ([]models.Transaction, error) {
	var transaction []models.Transaction
	if err := s.db.Where("user_id = ? AND title LIKE ?", userId, "%"+search+"%").Find(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return transaction, nil	
}

func (s *TransactionRepository) GetDetailByUniqueId(uniqueId string) (*models.Transaction, error) {
	var transaction models.Transaction
	query := s.db.Where("unique_id = ?", uniqueId)
	query = query.Preload("TransactionItems")
	query = query.Preload("TransactionItems.TransactionItemAddOns")
	query = query.Preload("TransactionItems.TransactionItemUsers")
	query = query.Preload("TransactionItems.TransactionItemUsers.User", func(db *gorm.DB) *gorm.DB {
        return db.Select("users.id", "users.name", "users.username", "users.email", "users.role_id", "users.is_active")
    })
	if err := query.First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

func (s *TransactionRepository) Create(transaction models.Transaction) (*models.Transaction, error) {
	if err := s.db.Create(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (s *TransactionRepository) Update(transaction models.Transaction) (*models.Transaction, error) {
	if err := s.db.Save(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}