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
	retrievedFolder, err := repo.getFolder(user.Username, folder1.Name)
	assert.NoError(t, err)
	assert.Contains(t, retrievedFolder.Files, &file1)

	// Scenario 4: Add a file with a duplicate name to a folder
	err = repo.CreateFile(user.Username, folder1.Name, file1)
	assert.Equal(t, errors.ResourceAlreadyExists(file1.Name), err)

	// Scenario 5: Check if files are sorted by name
	err = repo.CreateFile(user.Username, folder1.Name, file3)
	assert.NoError(t, err)
	err = repo.CreateFile(user.Username, folder1.Name, file2)
	assert.NoError(t, err)
	assert.Equal(t, []*entity.File{&file1, &file2, &file3}, retrievedFolder.Files)
}

func TestDeleteFile(t *testing.T) {
	var err error
	// Create a new in-memory repository
	repo := NewRepository()

	// Setup test user, folders and file
	user := entity.User{Username: "testuser"}
	folder1 := entity.Folder{Name: "testfolder1"}
	file1 := entity.File{Name: "testfile1", CreatedAt: time.Now()}
	file2 := entity.File{Name: "testfile2", CreatedAt: time.Now()}
	file3 := entity.File{Name: "testfile3", CreatedAt: time.Now()}

	// Scenario 1: Delete file from non-existent testuser
	err = repo.DeleteFile(user.Username, folder1.Name, file1.Name)
	assert.Equal(t, errors.ResourceNotFound("testuser"), err)

	err = repo.CreateUser(&user)
	assert.NoError(t, err)

	// Scenario 2: Delete file from non-existent folder
	err = repo.DeleteFile(user.Username, folder1.Name, file1.Name)
	assert.Equal(t, errors.ResourceNotFound("testfolder1"), err)

	err = repo.CreateFolder(user.Username, folder1)
	assert.NoError(t, err)

	// Scenario 3: Delete file that does not exist
	err = repo.DeleteFile(user.Username, folder1.Name, file1.Name)
	assert.Equal(t, errors.ResourceNotFound("testfile1"), err)

	// Scenario 4: Successfully delete file
	err = repo.CreateFile(user.Username, folder1.Name, file1)
	assert.NoError(t, err)
	err = repo.CreateFile(user.Username, folder1.Name, file2)
	assert.NoError(t, err)
	err = repo.CreateFile(user.Username, folder1.Name, file3)
	assert.NoError(t, err)
	err = repo.DeleteFile(user.Username, folder1.Name, file2.Name)
	assert.NoError(t, err)
	retrievedFolder, err := repo.getFolder(user.Username, folder1.Name)
	assert.NoError(t, err)
	assert.Equal(t, []*entity.File{&file1, &file3}, retrievedFolder.Files)
}

func TestListFiles(t *testing.T) {
	var err error
	// Create a new in-memory repository
	repo := NewRepository()

	// Setup test user, folders and file
	user := entity.User{Username: "testuser"}
	folder1 := entity.Folder{Name: "testfolder1", Files: []*entity.File{}}
	file1 := entity.File{Name: "testfile1", CreatedAt: time.Now().Add(-1 * time.Hour)}
	file2 := entity.File{Name: "testfile2", CreatedAt: time.Now()}
	file3 := entity.File{Name: "testfile3", CreatedAt: time.Now().Add(-2 * time.Hour)}
	err = repo.CreateUser(&user)
	assert.NoError(t, err)
	err = repo.CreateFolder(user.Username, folder1)
	assert.NoError(t, err)

	// Scenario 1: List files from non-existent testuser
	_, err = repo.ListFiles("non-existent", folder1.Name, entity.ListFileOption{})
	assert.Equal(t, errors.ResourceNotFound("non-existent"), err)

	// Scenario 2: List files from non-existent folder
	_, err = repo.ListFiles(user.Username, "non-existent", entity.ListFileOption{})
	assert.Equal(t, errors.ResourceNotFound("non-existent"), err)

	// Scenario 3: List files from an empty folder
	files, err := repo.ListFiles(user.Username, folder1.Name, entity.ListFileOption{})
	assert.NoError(t, err)
	assert.Equal(t, []*entity.File{}, files)

	// Scenario 4: List files from a folder with files
	err = repo.CreateFile(user.Username, folder1.Name, file1)
	assert.NoError(t, err)
	err = repo.CreateFile(user.Username, folder1.Name, file2)
	assert.NoError(t, err)
	err = repo.CreateFile(user.Username, folder1.Name, file3)
	assert.NoError(t, err)
	files, err = repo.ListFiles(user.Username, folder1.Name, entity.ListFileOption{})
	assert.NoError(t, err)
	assert.Equal(t, []*entity.File{&file1, &file2, &file3}, files)

	// Scenario 5: List files from a folder with files sorted by name
	files, err = repo.ListFiles(user.Username, folder1.Name, entity.ListFileOption{
		Sort: entity.SortOption{
			Attribute: entity.SortByName,
			Direction: entity.Desc,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []*entity.File{&file3, &file2, &file1}, files)

	// Scenario 6: List files from a folder with files sorted by create time
	files, err = repo.ListFiles(user.Username, folder1.Name, entity.ListFileOption{
		Sort: entity.SortOption{
			Attribute: entity.SortByCreateTime,
			Direction: entity.Asc,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, []*entity.File{&file3, &file1, &file2}, files)
}
