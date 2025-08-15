package main

import (
	"regexp"
	"strings"
)

func removeComments(content string) string {
	// Match comments: lines starting with #
	singleLineRegex := regexp.MustCompile(`^# `)

	commentedLines := regexp.MustCompile(`\r?\n`).Split(content, -1)

	uncommentedLines := make([]string, 0, len(commentedLines))

	for _, line := range commentedLines {
		newLine := line[:]
		if singleLineRegex.MatchString(newLine) {
			newLine = singleLineRegex.ReplaceAllString(newLine, "")
		}
		uncommentedLines = append(uncommentedLines, newLine)
	}

	return strings.Join(uncommentedLines, "\n")
}
