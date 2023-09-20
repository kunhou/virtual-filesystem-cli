package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
)

func Test_repository_GetUserByName(t *testing.T) {
	repo := NewRepository()

	_, err := repo.GetUserByName("test")
	assert.Error(t, err)
	assert.Equal(t, "The test doesn't exist.", err.Error())

	repo.users["test"] = &entity.User{
		Username: "test",
		Folders:  make([]*entity.Folder, 0),
	}

	_, err = repo.GetUserByName("test")
	assert.NoError(t, err)
}

func Test_repository_CreateUser(t *testing.T) {
	repo := NewRepository()

	err := repo.CreateUser(&entity.User{
		Username: "test",
	})
	assert.NoError(t, err)

	err = repo.CreateUser(&entity.User{
		Username: "test",
	})
	assert.Error(t, err)
	assert.Equal(t, "The test has already existed.", err.Error())
}
