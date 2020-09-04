package parse

import (
	"fmt"
	"strings"
	"unicode"
)

var labels = map[string]string{}

// Label parses a label
func Label(inp string, line uint16) (bool, error) {
	if match := Submatch(inp, `^([a-zA-Z0-9_]*):(?:\s+//.*)*$`); match != nil {
		if err := validateLabel(strings.ToLower(match[1])); err != nil {
			return false, err
		}
		labels[strings.ToLower(match[1])] = fmt.Sprintf("$%d", line)
		return true, nil
	}
	return false, nil
}

// GetLabel returns a stored label address
func GetLabel(label string) string {
	return labels[strings.ToLower(label)]
}

// ReplaceFromLabel returns the address of the label
// or returns the inputted string if not found
func ReplaceFromLabel(label string) string {
	if val := labels[strings.ToLower(label)]; val != "" {
		return val
	}
	return label
}

// ClearLabels clears all label addresses
// Mostly only used for testing and debuging purposes
func ClearLabels() {
	labels = map[string]string{}
}

func validateLabel(label string) error {
	if !unicode.IsLetter(rune(label[0])) {
		return fmt.Errorf("Label must start with a letter '%s'", label)
	}

	if labels[label] != "" {
		return fmt.Errorf("Duplicate label '%s'", label)
	}
	return nil

}
