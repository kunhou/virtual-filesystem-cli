package parser

import (
	"strings"
)

// ParseInput takes a raw input string and returns a slice of tokens.
// Tokens can be separated by spaces and can optionally be enclosed in double quotes.
func ParseInput(input string) []string {
	var tokens []string
	var currentToken strings.Builder
	inQuotes := false

	for _, c := range input {
		switch c {
		case ' ':
			if inQuotes {
				currentToken.WriteRune(c)
			} else {
				if currentToken.Len() > 0 {
					tokens = append(tokens, currentToken.String())
					currentToken.Reset()
				}
			}
		case '"':
			if inQuotes {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			inQuotes = !inQuotes
		default:
			currentToken.WriteRune(c)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
}
