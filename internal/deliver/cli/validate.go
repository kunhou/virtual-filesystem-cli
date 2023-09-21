package cli

import "regexp"

// validateName checks if the username adheres to established guidelines.
func validateName(name string) bool {
	// Check if the length is between 3 and 20 characters.
	if len(name) == 0 || len(name) > 20 {
		return false
	}

	// Check if it contains any invalid characters.
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", name)

	return matched
}
