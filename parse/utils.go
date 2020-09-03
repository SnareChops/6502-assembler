package parse

import "regexp"

// Submatch runs a regex on the string and returns all submatches
func Submatch(inp string, reg string) []string {
	regex := regexp.MustCompile(reg)
	return regex.FindStringSubmatch(inp)
}

// Matcher returns a matcher function
func Matcher(pattern string) func(string) []string {
	regex := regexp.MustCompile(pattern)
	return func(inp string) []string {
		return regex.FindStringSubmatch(inp)
	}
}
