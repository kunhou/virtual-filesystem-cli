package usecase

import (
	"time"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
)

func newFolder(name, description string) *entity.Folder {
	return &entity.Folder{
		Name:        name,
		Description: description,
		Files:       []*entity.File{},
		CreatedAt:   time.Now(),
	}
}

// CreateFolder creates a new folder.
func (u *Usecase) CreateFolder(opt entity.CreateFolderParam) error {
	folder := newFolder(opt.Name, opt.Description)
	return u.repo.CreateFolder(opt.Username, folder)
}

// DeleteFolder deletes a folder.
func (u *Usecase) DeleteFolder(username, folderName string) error {
	return u.repo.DeleteFolder(username, folderName)
}

// ListFolders lists all folders.
func (u *Usecase) ListFolders(username string, opt entity.ListFolderOption) ([]*entity.Folder, error) {
	return u.repo.ListFolders(username, opt)
}
