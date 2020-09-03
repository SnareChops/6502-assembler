package parse

import "regexp"

// Comment parses a comment
func Comment(inp string) bool {
	regex := regexp.MustCompile(`^//.*`)
	if match := regex.FindString(inp); match != "" {
		return true
	}
	return false
}
