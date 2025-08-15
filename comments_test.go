package main

import (
	"testing"
)

func TestRemoveComments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No comments",
			input:    "foo = \"bar\"\nbar = 123",
			expected: "foo = \"bar\"\nbar = 123",
		},
		{
			name:     "With comments",
			input:    "# # This is a comment\nfoo = \"bar\"\n# Another comment\nbar = 123",
			expected: "# This is a comment\nfoo = \"bar\"\nAnother comment\nbar = 123",
		},
		{
			name:     "Empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "Only comments",
			input:    "# Comment 1\n# Comment 2",
			expected: "Comment 1\nComment 2",
		},
		{
			name:     "Ignore different comment style",
			input:    "## Comment 1\n##Comment 2",
			expected: "## Comment 1\n##Comment 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeComments(tt.input)
			if result != tt.expected {
				t.Errorf("removeComments() = %q, want %q", result, tt.expected)
			}
		})
	}
}
