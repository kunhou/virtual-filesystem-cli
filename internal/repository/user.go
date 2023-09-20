package repository

import (
	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/pkg/errors"
)

// CreateUser adds a new user.
func (r *repository) CreateUser(user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Username]; exists {
		return errors.ResourceAlreadyExists(user.Username)
	}
	r.users[user.Username] = user
	return nil
}

// GetByUsername fetches a user by their username.
func (r *repository) GetUserByName(username string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.getUserByName(username)
}

// GetByUsername fetches a user by their username.
func (r *repository) getUserByName(username string) (*entity.User, error) {
	user, exists := r.users[username]
	if !exists {
		return nil, errors.ResourceNotFound(username)
	}
	return user, nil
}
