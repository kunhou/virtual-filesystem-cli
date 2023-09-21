package parser

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"simple test", []string{"simple", "test"}},
		{`"quoted string" here`, []string{"quoted string", "here"}},
		{`mixed "quoted string" and unquoted`, []string{"mixed", "quoted string", "and", "unquoted"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ParseInput(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
