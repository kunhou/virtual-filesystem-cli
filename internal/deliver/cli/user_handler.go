package cli

import (
	"regexp"

	"github/kunhou/virtual-filesystem-cli/pkg/log"
)

func (s *CLIServer) RegisterUserHandler(args []string) {
	if len(args) != 1 {
		log.Info("Usage: register [username]")
		return
	}

	username := args[0]

	if !validateUsername(username) {
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

// validateUsername checks if the username adheres to established guidelines.
func validateUsername(username string) bool {
	// Check if the length is between 3 and 20 characters.
	if len(username) < 3 || len(username) > 20 {
		return false
	}

	// Check if it contains any invalid characters.
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", username)

	return matched
}
