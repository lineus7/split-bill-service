package services

import (
	"split-bill-service/internal/models"
	"split-bill-service/internal/repositories"
)

func NewTransactionService(repos *repositories.RepositoriesSet) *TransactionService {
	return &TransactionService{
		transactionRepository: repos.TransactionRepository,
	}
}
type TransactionService struct {
	transactionRepository *repositories.TransactionRepository
}

func (s *TransactionService) GetListByUserId(userId uint) ([]models.Transaction, error) {
	return s.transactionRepository.GetListByUserId(userId)
}
