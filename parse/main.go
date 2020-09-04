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
		data = append(data, Line(scanner.Text(), uint16(len(data)))...)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

// Line parses a line of code
func Line(inp string, pointer uint16) []byte {
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

	if file := Include(inp); file != "" {
		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		bytes, err := File(f)
		if err != nil {
			panic(err)
		}
		f.Close()
		return bytes
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
	if inst, result := Either(inp, Instructions...); inst != "" {
		return result
	}

	panic(fmt.Errorf("Invalid line '%s'", inp))
}
