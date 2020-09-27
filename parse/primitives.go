package parse

import (
	"strconv"

	"github.com/snarechops/6502-assembler/conv"
)

// Uint16 parses a uint16 literal
func Uint16(inp string) (string, []byte) {
	i, err := strconv.ParseUint(inp, 0, 16)
	if err == nil {
		return "true", conv.Uint16(uint16(i))
	}
	return "", nil
}

// Int8 parses an int8 literal
func Int8(inp string) (string, []byte) {
	i, err := strconv.ParseInt(inp, 0, 8)
	if err == nil {
		return "true", []byte{byte(int8(i))}
	}
	return "", nil
}

// Uint8 parses a uint8 literal
func Uint8(inp string) (string, []byte) {
	i, err := strconv.ParseUint(inp, 0, 8)
	if err == nil {
		return "true", []byte{uint8(i)}
	}
	return "", nil
}

// Char parses a char literal
func Char(inp string) (string, []byte) {
	if match := Submatch(inp, `^'(.)'$`); match != nil {
		ascii := []rune(match[1])
		if ascii[0] < 256 {
			return "true", []byte{uint8(ascii[0])}
		}
	}
	return "", nil
}
