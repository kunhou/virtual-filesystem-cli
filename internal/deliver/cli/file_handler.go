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
