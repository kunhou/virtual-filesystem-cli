package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/pkg/errors"
)

func TestCreateFile(t *testing.T) {
	var err error
	// Create a new in-memory repository
	repo := NewRepository()

	// Setup test user and folders
	user := entity.User{Username: "testuser"}
	folder1 := entity.Folder{Name: "testfolder1"}

	// Dummy file
	file1 := entity.File{Name: "testfile1", CreatedAt: time.Now()}
	file2 := entity.File{Name: "testfile2", CreatedAt: time.Now()}
	file3 := entity.File{Name: "testfile3", CreatedAt: time.Now()}

	// Scenario 1: Add file to non-existent testuser
	err = repo.CreateFile(user.Username, folder1.Name, file1)
	assert.Equal(t, errors.ResourceNotFound("testuser"), err)

	err = repo.CreateUser(&user)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, folder1)
	assert.NoError(t, err)

	// Scenario 2: Add file to non-existent folder
	err = repo.CreateFile(user.Username, "non-existent-folder", file1)
	assert.Equal(t, errors.ResourceNotFound("non-existent-folder"), err)

	// Scenario 3: Successfully add file to an existing folder
	err = repo.CreateFile(user.Username, folder1.Name, file1)
	assert.NoError(t, err)
	retrievedUser, err := repo.GetUserByName(user.Username)
	assert.NoError(t, err)
	assert.Contains(t, retrievedUser.Folders[0].Files, &file1)

	// Scenario 4: Add a file with a duplicate name to a folder
	err = repo.CreateFile(user.Username, folder1.Name, file1)
	assert.Equal(t, errors.ResourceAlreadyExists(file1.Name), err)

	// Scenario 5: Check if files are sorted by name
	err = repo.CreateFile(user.Username, folder1.Name, file3)
	assert.NoError(t, err)
	err = repo.CreateFile(user.Username, folder1.Name, file2)
	assert.NoError(t, err)
	retrievedUser, err = repo.GetUserByName(user.Username)
	assert.NoError(t, err)
	assert.Equal(t, []*entity.File{&file1, &file2, &file3}, retrievedUser.Folders[0].Files)
}
