package parse

import "regexp"

// Submatch runs a regex on the string and returns all submatches
func Submatch(inp string, reg string) []string {
	regex := regexp.MustCompile(reg)
	return regex.FindStringSubmatch(inp)
}
