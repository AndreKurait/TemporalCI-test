package main

import "strings"

// TitleCase converts a string to title case.
func TitleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

// Truncate shortens a string to maxLen, adding "..." if truncated.
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// CountWords returns the number of words in a string.
func CountWords(s string) int {
	return len(strings.Fields(s))
}
