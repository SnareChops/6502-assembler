package parse

import (
	"strconv"

	"github.com/snarechops/assembler/conv"
)

// PrimitiveParser the type of a primitive parser
type PrimitiveParser = func(string) (bool, []byte)

// Uint16 parses a uint16 literal
func Uint16(inp string) (bool, []byte) {
	i, err := strconv.ParseUint(inp, 0, 16)
	if err == nil {
		return true, conv.Uint16(uint16(i))
	}
	return false, nil
}

// Uint8 parses a uint8 literal
func Uint8(inp string) (bool, []byte) {
	i, err := strconv.ParseUint(inp, 0, 8)
	if err == nil {
		return true, []byte{uint8(i)}
	}
	return false, nil
}

// Char parses a char literal
func Char(inp string) (bool, []byte) {
	if match := Submatch(inp, `^'(.)'$`); match != nil {
		ascii := []rune(match[1])
		if ascii[0] < 256 {
			return true, []byte{uint8(ascii[0])}
		}
	}
	return false, nil
}
