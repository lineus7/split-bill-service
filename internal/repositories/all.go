package repositories

import (
	"split-bill-service/config"
)

type RepositoriesSet struct {
	UserRepository *UserRepository
	GeminiRepository *GeminiRepository
	TransactionRepository *TransactionRepository
}

func NewRepositoriesSet(connection *config.Connection) *RepositoriesSet {
	return &RepositoriesSet{
		UserRepository: NewUserRepository(connection),
		GeminiRepository: NewGeminiRepository(connection),
		TransactionRepository: NewTransactionRepository(connection),
	}
}
