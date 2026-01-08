package utils

import "strings"

// Contains checks if a string contains a substring
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// HasPrefix checks if a string starts with a prefix
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// Join joins a slice of strings with a separator
func Join(elements []string, separator string) string {
	return strings.Join(elements, separator)
}

// Split splits a string by a separator
func Split(s, separator string) []string {
	return strings.Split(s, separator)
}

// TrimSpace removes leading and trailing whitespace
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}