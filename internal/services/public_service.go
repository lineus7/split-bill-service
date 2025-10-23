package services

import (
	"net/http"
	"split-bill-service/internal/models"
	"split-bill-service/internal/repositories"
	"split-bill-service/utils"
)

func NewPublicService(repos *repositories.RepositoriesSet) *PublicService {
	return &PublicService{
		transactionRepository: repos.TransactionRepository,
	}
}
type PublicService struct {
	transactionRepository *repositories.TransactionRepository
}

func (s *PublicService) GetTransactionDetailByUniqueId(uniqueId string) (*models.Transaction, error) {
	transaction, err := s.transactionRepository.GetDetailByUniqueId(uniqueId)
	if err != nil {
		return nil, err
	}
	if transaction == nil {
		return nil, utils.NewAppError(http.StatusNotFound, "Transaction not found", nil)
	}
	return transaction, nil
}
