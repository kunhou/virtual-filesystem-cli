package cli

import (
	"github/kunhou/virtual-filesystem-cli/pkg/log"
)

func (s *CLIServer) RegisterUserHandler(args []string) {
	if len(args) != 1 {
		log.Info("Usage: register [username]")
		return
	}

	username := args[0]

	if !validateName(username) {
		log.Error("The %s contain invalid chars.", username)
		return
	}

	err := s.usecase.CreateUser(username)
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Add %s successfully.", username)
}
