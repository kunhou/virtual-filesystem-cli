package errors

import "testing"

func TestError_Error(t *testing.T) {
	tests := []struct {
		template string
		args     []string
		expected string
	}{
		{
			template: "The %s doesn't exist.",
			args:     []string{"username"},
			expected: "The username doesn't exist.",
		},
		{
			template: "The %s has already existed.",
			args:     []string{"foldername"},
			expected: "The foldername has already existed.",
		},
		{
			template: "The %s / %s has already existed.",
			args:     []string{"foldername", "filename"},
			expected: "The foldername / filename has already existed.",
		},
	}

	for _, tt := range tests {
		err := &Error{
			messageTemplate: tt.template,
			messageArgs:     tt.args,
		}
		if err.Error() != tt.expected {
			t.Errorf("Expected %q, but got %q", tt.expected, err.Error())
		}
	}
}
