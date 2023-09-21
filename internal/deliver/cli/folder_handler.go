package cli

import (
	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/pkg/log"
)

func (s *CLIServer) CreateFolderHandler(args []string) {
	if len(args) < 2 {
		log.Error("Usage: create-folder [username] [foldername] [description]?")
		return
	}

	username := args[0]
	folderName := args[1]
	var description string
	if len(args) > 2 {
		description = args[2]
	}

	if !validateName(folderName) {
		log.Error("The %s contain invalid chars.", folderName)
		return
	}

	err := s.usecase.CreateFolder(entity.CreateFolderParam{
		Username:    username,
		Name:        folderName,
		Description: description,
	})
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Create %s successfully.", folderName)
}

func (s *CLIServer) DeleteFolderHandler(args []string) {
	if len(args) != 2 {
		log.Error("Usage: delete-folder [username] [foldername]")
		return
	}

	username := args[0]
	folderName := args[1]

	err := s.usecase.DeleteFolder(username, folderName)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Delete %s successfully.", folderName)
}

func (s *CLIServer) ListFoldersHandler(args []string) {
	if len(args) < 1 {
		log.Error("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
		return
	}

	username := args[0]

	attribute, direction, err := argsToSortOptions(args)
	if err != nil {
		log.Error(err.Error())
		return
	}

	folders, err := s.usecase.ListFolders(username, entity.ListFolderOption{
		Sort: entity.SortOption{
			Attribute: attribute,
			Direction: direction,
		},
	})
	if err != nil {
		log.Error(err.Error())
		return
	}

	if len(folders) == 0 {
		log.Warn("The %s doesn't have any folders.", username)
		return
	}

	for _, folder := range folders {
		log.Info("%s\t%s\t%s\t%s", folder.Name, folder.Description, folder.CreatedAt.Format("2006-01-02 15:04:05"), username)
	}
}
