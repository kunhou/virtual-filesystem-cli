package cli

import (
	"fmt"
	"regexp"
	"strings"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
)

const (
	MaxNameLength        = 20
	MaxDescriptionLength = 255
)

var (
	ErrInvalidSortName = fmt.Errorf("Invalid sort option. Use --sort-name or --sort-created.")
	ErrInvalidSortDir  = fmt.Errorf("Invalid sort direction. Use asc or desc.")
)

// validateName checks if the username adheres to established guidelines.
func validateName(name string) bool {
	// Check if the length is between 1 and 20 characters.
	if len(name) == 0 || len(name) > MaxNameLength {
		return false
	}

	// Check if it contains any invalid characters.
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", name)

	return matched
}

func validDescription(description string) bool {
	// Check if the length is between 1 and 255 characters.
	if len(description) == 0 || len(description) > MaxDescriptionLength {
		return false
	}

	return true
}

func argsToSortOptions(args []string) (attribute entity.SortAttribute, direction entity.SortDirection, err error) {
	attribute = entity.SortByName
	direction = entity.Asc

	if len(args) > 0 {
		switch args[0] {
		case "--sort-name":
			attribute = entity.SortByName
		case "--sort-created":
			attribute = entity.SortByCreateTime
		default:
			return attribute, direction, ErrInvalidSortName
		}
	}

	if len(args) > 1 {
		switch strings.ToLower(args[1]) {
		case "asc":
			direction = entity.Asc
		case "desc":
			direction = entity.Desc
		default:
			return attribute, direction, ErrInvalidSortDir
		}
	}

	return attribute, direction, nil
}
