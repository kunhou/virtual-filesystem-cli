package usecase

import (
	"github/kunhou/virtual-filesystem-cli/internal/entity"
)

func newUser(username string) *entity.User {
	return &entity.User{
		Username: username,
		Folders:  []*entity.Folder{},
	}
}

// CreateUser creates a new user.
func (u *Usecase) CreateUser(username string) error {
	user := newUser(username)
	return u.repo.CreateUser(user)
}
