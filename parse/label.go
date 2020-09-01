package parse

import (
	"fmt"
	"strings"

	"github.com/snarechops/assembler/conv"
)

var labels = map[string][]byte{}

// Label parses a label
func Label(inp string, line uint16) (bool, error) {
	if match := Submatch(inp, `^([a-zA-Z0-9_]*):$`); match != nil {
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

func validateLabel(label string) error {
	if labels[label] != nil {
		return fmt.Errorf("Duplicate label '%s'", label)
	}
	return nil
}
