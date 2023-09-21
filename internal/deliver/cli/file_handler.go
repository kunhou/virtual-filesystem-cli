package cli

import (
	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/pkg/log"
)

func (s *CLIServer) CreateFileHandler(args []string) {
	if len(args) < 3 {
		log.Error("Usage: create-file [username] [foldername] [filename] [description]?")
		return
	}

	username := args[0]
	folderName := args[1]
	fileName := args[2]

	var description string
	if len(args) > 3 {
		description = args[3]
	}

	err := s.usecase.CreateFile(entity.CreateFileParam{
		Username:    username,
		FolderName:  folderName,
		Name:        fileName,
		Description: description,
	})
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Create %s in %s / %s successfully.", fileName, username, folderName)
}

func (s *CLIServer) DeleteFileHandler(args []string) {
	if len(args) != 3 {
		log.Error("Usage: delete-file [username] [foldername] [filename]")
		return
	}

	username := args[0]
	folderName := args[1]
	fileName := args[2]

	err := s.usecase.DeleteFile(username, folderName, fileName)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Delete %s in %s / %s successfully.", fileName, username, folderName)
}

func (s *CLIServer) ListFilesHandler(args []string) {
	if len(args) < 2 {
		log.Error("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
		return
	}

	username := args[0]
	folderName := args[1]

	attribute, direction, err := argsToSortOptions(args[2:])
	if err != nil {
		log.Error(err.Error())
		return
	}

	files, err := s.usecase.ListFiles(username, folderName, entity.ListFileOption{
		Sort: entity.SortOption{
			Attribute: attribute,
			Direction: direction,
		},
	})
	if err != nil {
		log.Error(err.Error())
		return
	}

	if len(files) == 0 {
		log.Warn("The folder is empty.")
		return
	}

	for _, file := range files {
		log.Info("%s\t%s\t%s\t%s\t%s", file.Name, file.Description, file.CreatedAt.Format(entity.ListResourceTimeFormat), folderName, username)
	}
}
