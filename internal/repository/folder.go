package repository

import (
	"sort"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/pkg/errors"
)

// CreateFolder adds a new folder to a user.
func (r *repository) CreateFolder(username string, folder *entity.Folder) error {
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

	user.Folders = append(user.Folders, folder)

	// Sort folders by name after adding a new one.
	sort.Slice(user.Folders, func(i, j int) bool {
		return user.Folders[i].Name < user.Folders[j].Name
	})

	r.users[username] = user
	return nil
}

// DeleteFolder removes a folder from a user.
func (r *repository) DeleteFolder(username, folderName string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the user exists.
	user, err := r.getUserByName(username)
	if err != nil {
		return err
	}

	for i, f := range user.Folders {
		if f.Name == folderName {
			user.Folders = append(user.Folders[:i], user.Folders[i+1:]...)
			r.users[username] = user // Update the map with the modified user.
			return nil
		}
	}

	return errors.ResourceNotFound(folderName)
}

func (r *repository) ListFolders(username string, opt entity.ListFolderOption) ([]*entity.Folder, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, err := r.getUserByName(username)
	if err != nil {
		return nil, err
	}

	folders := user.Folders
	sort.Slice(folders, func(i, j int) bool {
		return sortFoldersByAttribute(opt.Sort.Attribute, opt.Sort.Direction)(folders[i], folders[j])
	})

	return folders, nil
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

func sortFoldersByAttribute(attribute entity.SortAttribute, direction entity.SortDirection) func(fi, fj *entity.Folder) bool {
	switch {
	case attribute == entity.SortByCreateTime && direction == entity.Asc:
		return func(fi, fj *entity.Folder) bool {
			return fi.CreatedAt.Before(fj.CreatedAt)
		}
	case attribute == entity.SortByCreateTime && direction == entity.Desc:
		return func(fi, fj *entity.Folder) bool {
			return fi.CreatedAt.After(fj.CreatedAt)
		}
	case attribute == entity.SortByName && direction == entity.Asc:
		return func(fi, fj *entity.Folder) bool {
			return fi.Name < fj.Name
		}
	case attribute == entity.SortByName && direction == entity.Desc:
		return func(fi, fj *entity.Folder) bool {
			return fi.Name > fj.Name
		}
	default:
		return func(fi, fj *entity.Folder) bool {
			return fi.Name < fj.Name
		}
	}
}
