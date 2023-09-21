package repository

import (
	"sort"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/pkg/errors"
)

func (r *repository) CreateFile(username string, folderName string, file entity.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	folder, err := r.getFolder(username, folderName)
	if err != nil {
		return err
	}

	for _, f := range folder.Files {
		if f.Name == file.Name {
			return errors.ResourceAlreadyExists(file.Name)
		}
	}

	folder.Files = append(folder.Files, &file)

	// Sort folders by name after adding a new one.
	sort.Slice(folder.Files, func(i, j int) bool {
		return folder.Files[i].Name < folder.Files[j].Name
	})

	return nil
}

func (r *repository) DeleteFile(username, folderName, fileName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	folder, err := r.getFolder(username, folderName)
	if err != nil {
		return err
	}

	for i, f := range folder.Files {
		if f.Name == fileName {
			folder.Files = append(folder.Files[:i], folder.Files[i+1:]...)
			return nil
		}
	}

	return errors.ResourceNotFound(fileName)
}
