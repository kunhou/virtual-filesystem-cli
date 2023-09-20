package usecase

import (
	"github/kunhou/virtual-filesystem-cli/internal/entity"
)

type IRepository interface {
	CreateUser(user *entity.User) error
	GetUserByName(username string) (*entity.User, error)

	CreateFolder(username string, folder entity.Folder) error
	DeleteFolder(username, folderName string) error
	ListFolders(username string, opt entity.ListFolderOption) ([]*entity.Folder, error)
}
