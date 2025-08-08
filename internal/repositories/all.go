package repositories

import (
	"split-bill-service/config"
)

type RepositoriesSet struct {
	UserRepository *UserRepository
}

func NewRepositoriesSet(connection *config.Connection) *RepositoriesSet {
	return &RepositoriesSet{
		UserRepository: NewUserRepository(connection),
	}
}
