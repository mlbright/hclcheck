package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclparse"
)

func main() {
	// Define command-line flags
	var verbose bool

	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")

	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path-to-tfvars-file>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := args[0]

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: File '%s' does not exist\n", filePath)
		os.Exit(1)
	}

	if err := findErrors(filePath, verbose); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing file '%s': %v\n", filePath, err)
	}
}

func findErrors(filePath string, verbose bool) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file: %w\n", err)
	}

	uncommented := removeComments(string(content))

	parser := hclparse.NewParser()

	_, diags := parser.ParseHCL([]byte(uncommented), filePath)

	if diags.HasErrors() {
		fmt.Println("detected invalid .tfvars file:", filePath)
		if verbose {
			for _, diag := range diags {
				fmt.Printf("- %s", diag.Error())
			}
		}
	}

	return nil
}
