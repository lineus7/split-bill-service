package repositories

import (
	"split-bill-service/config"
)

type RepositoriesSet struct {
	UserRepository *UserRepository
	GeminiRepository *GeminiRepository
}

func NewRepositoriesSet(connection *config.Connection) *RepositoriesSet {
	return &RepositoriesSet{
		UserRepository: NewUserRepository(connection),
		GeminiRepository: NewGeminiRepository(connection),
	}
}
