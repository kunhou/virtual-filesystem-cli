package repository

import (
	"sort"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/pkg/errors"
)

// CreateFolder adds a new folder to a user.
func (r *repository) CreateFolder(username string, folder entity.Folder) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the user exists.
	user, err := r.getUserByName(username)
	if err != nil {
		return err
	}

	for _, f := range user.Folders {
		if f.Name == folder.Name {
			return errors.ResourceAlreadyExists(folder.Name)
		}
	}

	user.Folders = append(user.Folders, &folder)

	// Sort folders by name after adding a new one.
	sort.Slice(user.Folders, func(i, j int) bool {
		return user.Folders[i].Name < user.Folders[j].Name
	})

	r.users[username] = user
	return nil
}


func (r *repository) getFolder(username, folderName string) (*entity.Folder, error) {
	// Check if the user exists.
	user, err := r.getUserByName(username)
	if err != nil {
		return nil, err
	}

	for _, f := range user.Folders {
		if f.Name == folderName {
			return f, nil
		}
	}
	return nil, errors.ResourceNotFound(folderName)
}
