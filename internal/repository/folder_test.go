package repository

import (
	"testing"

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
	err := repo.CreateFolder(user.Username, folder1)
	assert.Equal(t, errors.ResourceNotFound("testuser"), err)

	// Scenario 2: Successfully add folder to an existing user
	err = repo.CreateUser(&user)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, folder1)
	assert.NoError(t, err)
	retrievedUser, err := repo.GetUserByName(user.Username)
	assert.NoError(t, err)
	assert.Contains(t, retrievedUser.Folders, &folder1)

	// Scenario 3: Add a folder with a duplicate name to a user
	err = repo.CreateFolder(user.Username, folder1)
	assert.Equal(t, errors.ResourceAlreadyExists(folder1.Name), err)

	// Scenario 4: Check if folders are sorted by name
	folder0 := entity.Folder{Name: "testfolder0"}
	err = repo.CreateFolder(user.Username, folder0)
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
	err = repo.CreateFolder(user.Username, folder1)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, folder2)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, folder3)
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
	err = repo.CreateFolder(user.Username, folder1)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, folder2)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, folder3)
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
