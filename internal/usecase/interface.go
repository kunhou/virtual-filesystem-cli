package usecase

import (
	"github/kunhou/virtual-filesystem-cli/internal/entity"
)

type IRepository interface {
	CreateUser(user *entity.User) error
	GetUserByName(username string) (*entity.User, error)
}
