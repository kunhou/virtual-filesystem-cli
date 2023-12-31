package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/pkg/errors"
)

func TestCreateFolder(t *testing.T) {
	// Create a new in-memory repository
	repo := NewRepository()

	// Dummy user and folder
	user := entity.User{Username: "testuser"}
	folder1 := entity.Folder{Name: "testfolder1"}

	// Scenario 1: Add folder to non-existent testuser
	err := repo.CreateFolder(user.Username, &folder1)
	assert.Equal(t, errors.ResourceNotFound("testuser"), err)

	// Scenario 2: Successfully add folder to an existing user
	err = repo.CreateUser(&user)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder1)
	assert.NoError(t, err)
	retrievedUser, err := repo.GetUserByName(user.Username)
	assert.NoError(t, err)
	assert.Contains(t, retrievedUser.Folders, &folder1)

	// Scenario 3: Add a folder with a duplicate name to a user
	err = repo.CreateFolder(user.Username, &folder1)
	assert.Equal(t, errors.ResourceAlreadyExists(folder1.Name), err)

	// Scenario 4: Check if folders are sorted by name
	folder0 := entity.Folder{Name: "testfolder0"}
	err = repo.CreateFolder(user.Username, &folder0)
	assert.NoError(t, err)
	retrievedUser, err = repo.GetUserByName(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, []*entity.Folder{&folder0, &folder1}, retrievedUser.Folders)
}

func Test_getFolder(t *testing.T) {
	// Create a new in-memory repository
	repo := NewRepository()

	// Setup test user and folders
	user := entity.User{Username: "testuser"}
	folder1 := entity.Folder{Name: "testfolder1"}
	folder2 := entity.Folder{Name: "testfolder2"}
	folder3 := entity.Folder{Name: "testfolder3"}
	var err error
	err = repo.CreateUser(&user)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder1)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder2)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder3)
	assert.NoError(t, err)

	// Scenario 1: Get a folder that exists
	retrievedFolder, err := repo.getFolder(user.Username, folder2.Name)
	assert.NoError(t, err)
	assert.Equal(t, &folder2, retrievedFolder)

	// Scenario 2: Get a folder that does not exist
	retrievedFolder, err = repo.getFolder(user.Username, "nonexistentfolder")
	assert.Equal(t, errors.ResourceNotFound("nonexistentfolder"), err)
	assert.Nil(t, retrievedFolder)

	// Scenario 3: Get a folder from a user that does not exist
	retrievedFolder, err = repo.getFolder("nonexistentuser", folder1.Name)
	assert.Equal(t, errors.ResourceNotFound("nonexistentuser"), err)
	assert.Nil(t, retrievedFolder)
}

func TestDeleteFolder(t *testing.T) {
	// Create a new in-memory repository
	repo := NewRepository()

	// Setup test user and folders
	user := entity.User{Username: "testuser"}
	folder1 := entity.Folder{Name: "testfolder1"}
	folder2 := entity.Folder{Name: "testfolder2"}
	folder3 := entity.Folder{Name: "testfolder3"}
	var err error
	err = repo.CreateUser(&user)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder1)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder2)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder3)
	assert.NoError(t, err)

	// Scenario 1: Delete a folder that exists
	err = repo.DeleteFolder(user.Username, folder2.Name)
	assert.NoError(t, err)
	retrievedUser, err := repo.GetUserByName(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, []*entity.Folder{&folder1, &folder3}, retrievedUser.Folders)

	// Scenario 2: Delete a folder that does not exist
	err = repo.DeleteFolder(user.Username, "nonexistentfolder")
	assert.Equal(t, errors.ResourceNotFound("nonexistentfolder"), err)

	// Scenario 3: Delete a folder from a user that does not exist
	err = repo.DeleteFolder("nonexistentuser", folder1.Name)
	assert.Equal(t, errors.ResourceNotFound("nonexistentuser"), err)
}

func TestListFolders(t *testing.T) {
	// Create a new in-memory repository
	repo := NewRepository()

	// Setup test user and folders
	user := entity.User{Username: "testuser"}
	folder1 := entity.Folder{Name: "testfolder1", CreatedAt: time.Now().Add(-1 * time.Hour)}
	folder2 := entity.Folder{Name: "testfolder2", CreatedAt: time.Now()}
	folder3 := entity.Folder{Name: "testfolder3", CreatedAt: time.Now().Add(-2 * time.Hour)}
	var err error
	err = repo.CreateUser(&user)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder1)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder2)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder3)
	assert.NoError(t, err)

	// Scenario 1: List folders in ascending order by name
	folders, err := repo.ListFolders(user.Username, entity.ListFolderOption{
		Sort: entity.SortOption{
			Attribute: entity.SortByName,
			Direction: entity.Asc,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []*entity.Folder{&folder1, &folder2, &folder3}, folders)

	// Scenario 2: List folders in descending order by name
	folders, err = repo.ListFolders(user.Username, entity.ListFolderOption{
		Sort: entity.SortOption{
			Attribute: entity.SortByName,
			Direction: entity.Desc,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []*entity.Folder{&folder3, &folder2, &folder1}, folders)

	// Scenario 3: List folders in ascending order by create time
	folders, err = repo.ListFolders(user.Username, entity.ListFolderOption{
		Sort: entity.SortOption{
			Attribute: entity.SortByCreateTime,
			Direction: entity.Asc,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []*entity.Folder{&folder3, &folder1, &folder2}, folders)

	// Scenario 4: List folders in descending order by create time
	folders, err = repo.ListFolders(user.Username, entity.ListFolderOption{
		Sort: entity.SortOption{
			Attribute: entity.SortByCreateTime,
			Direction: entity.Desc,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []*entity.Folder{&folder2, &folder1, &folder3}, folders)

	// Scenario 5: List folders from a user that does not exist
	folders, err = repo.ListFolders("nonexistentuser", entity.ListFolderOption{})
	assert.Equal(t, errors.ResourceNotFound("nonexistentuser"), err)
	assert.Nil(t, folders)
}

func TestRenameFolder(t *testing.T) {
	// Create a new in-memory repository
	repo := NewRepository()

	// Setup test user and folders
	user := entity.User{Username: "testuser"}
	folder1 := entity.Folder{Name: "testfolder1", CreatedAt: time.Now().Add(-1 * time.Hour)}
	folder2 := entity.Folder{Name: "testfolder2", CreatedAt: time.Now()}
	folder3 := entity.Folder{Name: "testfolder3", CreatedAt: time.Now().Add(-2 * time.Hour)}
	var err error
	err = repo.CreateUser(&user)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder1)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder2)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, &folder3)
	assert.NoError(t, err)

	// Scenario 1: Rename a folder that exists
	err = repo.RenameFolder(user.Username, folder2.Name, "newfoldername")
	assert.NoError(t, err)
	folder, err := repo.getFolder(user.Username, "newfoldername")
	assert.NoError(t, err)
	assert.Equal(t, "newfoldername", folder.Name)

	// Scenario 2: Rename a folder that does not exist
	err = repo.RenameFolder(user.Username, "nonexistentfolder", "newfoldername2")
	assert.Equal(t, errors.ResourceNotFound("nonexistentfolder"), err)

	// Scenario 3: Rename a folder from a user that does not exist
	err = repo.RenameFolder("nonexistentuser", folder1.Name, "newfoldername")
	assert.Equal(t, errors.ResourceNotFound("nonexistentuser"), err)

	// Scenario 4: Rename a folder to a name that already exists
	err = repo.RenameFolder(user.Username, folder1.Name, folder2.Name)
	assert.Equal(t, errors.ResourceAlreadyExists(folder2.Name), err)

	// Scenario 5: Rename a folder to the same name
	err = repo.RenameFolder(user.Username, folder1.Name, folder1.Name)
	assert.NoError(t, err)
	folder, err = repo.getFolder(user.Username, folder1.Name)
	assert.NoError(t, err)
	assert.Equal(t, folder1.Name, folder.Name)
}
