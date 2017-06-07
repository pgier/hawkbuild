package main

import "strings"

// StringInSlice checks if a string is contained in the slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IsEmptyString checks if a string contains only whitespace
func IsEmptyString(aStr string) bool {
	return len(strings.TrimSpace(aStr)) == 0
}
