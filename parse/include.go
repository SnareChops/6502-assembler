package parse

// Include parses an include statement
func Include(inp string) string {
	if match := Submatch(inp, `(?i)^\.include\s+"?([^\n\r"]*)"?$`); match != nil {
		return match[1]
	}
	return ""
}
