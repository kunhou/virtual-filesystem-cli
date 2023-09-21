package cli

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{"validName", true},
		{"", false},
		{"excessivel yongnamess", false},
		{"name-with_special!char", false},
		{"another-valid_name", true},
		{"12 45lder-name", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, validateName(tt.name))
		})
	}
}

func TestValidDescription(t *testing.T) {
	tests := []struct {
		description string
		expected    bool
	}{
		{"This is a valid description.", true},
		{"", false}, // Empty string
		{"a", true}, // Smallest non-empty string
		{strings.Repeat("a", MaxDescriptionLength), true},
		{strings.Repeat("a", MaxDescriptionLength+1), false},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			assert.Equal(t, tt.expected, validDescription(tt.description))
		})
	}
}

func TestArgsToSortOptions(t *testing.T) {
	tests := []struct {
		args              []string
		expectedAttribute entity.SortAttribute
		expectedDirection entity.SortDirection
		expectedErr       error
	}{
		{[]string{}, entity.SortByName, entity.Asc, nil},
		{[]string{"--sort-name"}, entity.SortByName, entity.Asc, nil},
		{[]string{"--sort-created"}, entity.SortByCreateTime, entity.Asc, nil},
		{[]string{"--sort-name", "asc"}, entity.SortByName, entity.Asc, nil},
		{[]string{"--sort-name", "desc"}, entity.SortByName, entity.Desc, nil},
		{[]string{"--sort-invalid"}, entity.SortByName, entity.Asc, ErrInvalidSortParam},
		{[]string{"--sort-name", "invalid"}, entity.SortByName, entity.Asc, ErrInvalidSortParam},
	}

	for _, tt := range tests {
		t.Run("args:"+strings.Join(tt.args, "|"), func(t *testing.T) {
			attribute, direction, err := argsToSortOptions(tt.args)
			assert.Equal(t, tt.expectedAttribute, attribute)
			assert.Equal(t, tt.expectedDirection, direction)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
