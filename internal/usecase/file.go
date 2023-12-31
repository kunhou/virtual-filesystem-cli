package usecase

import (
	"time"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
)

func newFile(name, description string) *entity.File {
	return &entity.File{
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

func (u *Usecase) CreateFile(opt entity.CreateFileParam) error {
	file := newFile(opt.Name, opt.Description)
	return u.repo.CreateFile(opt.Username, opt.FolderName, file)
}

func (u *Usecase) DeleteFile(username, folderName, fileName string) error {
	return u.repo.DeleteFile(username, folderName, fileName)
}

func (u *Usecase) ListFiles(username, folderName string, opt entity.ListFileOption) ([]*entity.File, error) {
	return u.repo.ListFiles(username, folderName, opt)
}
