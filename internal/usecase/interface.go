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

	CreateFile(username string, folderName string, file entity.File) error
	DeleteFile(username, folderName, fileName string) error
	ListFiles(username, folderName string, opt entity.ListFileOption) ([]*entity.File, error)
}

type Usecase struct {
	repo IRepository
}

func NewUsecase(repo IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}
