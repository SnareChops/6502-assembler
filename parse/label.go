package parse

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/snarechops/assembler/conv"
)

var labels = map[string][]byte{}

// Label parses a label
func Label(inp string, line uint16) (bool, error) {
	if match := Submatch(inp, `^([a-zA-Z0-9_]*):(?:\s+//.*)*$`); match != nil {
		if err := validateLabel(strings.ToLower(match[1])); err != nil {
			return false, err
		}
		labels[strings.ToLower(match[1])] = conv.Uint16(line)
		return true, nil
	}
	return false, nil
}

// GetLabel returns a stored label address
func GetLabel(label string) []byte {
	return labels[strings.ToLower(label)]
}

// ClearLabels clears all label addresses
// Mostly only used for testing and debuging purposes
func ClearLabels() {
	labels = map[string][]byte{}
}

func validateLabel(label string) error {
	if !unicode.IsLetter(rune(label[0])) {
		return fmt.Errorf("Label must start with a letter '%s'", label)
	}

	if labels[label] != nil {
		return fmt.Errorf("Duplicate label '%s'", label)
	}
	return nil

}
