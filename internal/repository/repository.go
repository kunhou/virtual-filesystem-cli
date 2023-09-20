package repository

import (
	"sync"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/internal/usecase"
)

var _ usecase.IRepository = (*repository)(nil)

// repository implements repository interface
type repository struct {
	mu sync.RWMutex

	users map[string]*entity.User
}

// NewRepository returns new repository
func NewRepository() *repository {
	return &repository{
		users: make(map[string]*entity.User),
	}
}
