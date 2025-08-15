package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFindErrors(t *testing.T) {
	tests := []struct {
		name        string
		fileContent string
		expectError bool
		verbose     bool
	}{
		{
			name: "Valid commenting",
			fileContent: `
# # This is a valid comment describing the string variable
# variable = "value"
# # This is a valid comment describing the number variable
# number = 42
# # This is a valid comment describing the list variable
# list = ["item1", "item2"]
`,
			expectError: false,
			verbose:     false,
		},
		{
			name: "Invalid commenting",
			fileContent: `
# # This is a valid comment describing the string variable
# variable = "value"
# This is an INVALID comment describing the number variable
# number = 42,
# # This is a valid comment describing the list variable
# list = ["item1", "item2"]
`,
			expectError: true,
			verbose:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file
			tempDir, err := ioutil.TempDir("", "tfvars-test")
			if err != nil {
				t.Fatalf("could not create temp dir: %v", err)
			}
			defer os.RemoveAll(tempDir)

			filePath := filepath.Join(tempDir, "test.tfvars")
			if err := ioutil.WriteFile(filePath, []byte(tt.fileContent), 0644); err != nil {
				t.Fatalf("could not write temp file: %v", err)
			}

			// Capture stdout to check output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call the function
			err = findErrors(filePath, tt.verbose)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout
			out, _ := ioutil.ReadAll(r)
			output := string(out)

			// Check error
			if err != nil {
				t.Errorf("findErrors() returned error: %v", err)
			}

			// Check if error was detected as expected
			if tt.expectError && !strings.Contains(output, "detected invalid .tfvars file") {
				t.Errorf("expected error detection, but none found in output: %s", output)
			}

			if !tt.expectError && strings.Contains(output, "detected invalid .tfvars file") {
				t.Errorf("unexpected error detection in output: %s", output)
			}

			// Check verbose output
			if tt.verbose && tt.expectError && !strings.Contains(output, "-") {
				t.Errorf("verbose mode enabled but no detailed errors shown in output: %s", output)
			}
		})
	}
}
