package parse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// File parses a file
func File(file *os.File) ([]byte, error) {
	scanner := bufio.NewScanner(file)
	data := []byte{}
	for scanner.Scan() {
		data = append(data, Line(scanner.Text())...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

var pointer uint16 = 0

// Line parses a line of code
func Line(inp string) []byte {
	// Trim the string
	inp = strings.TrimSpace(inp)

	// Skip empty lines
	if len(inp) == 0 {
		return []byte{}
	}

	// Skip comments
	if Comment(inp) {
		return []byte{}
	}

	// Parse any labels
	match, err := Label(inp, pointer)
	if err != nil {
		panic(err)
	}
	if match {
		return []byte{}
	}

	// Parse any instructions
	if inst, result := Either(inp, LDA, LDX, LDY, STA, STX, STY); inst != "" {
		pointer++
		return result
	}

	panic(fmt.Errorf("Invalid line '%s'", inp))
}
