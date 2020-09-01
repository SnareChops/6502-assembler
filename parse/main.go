package parse

import (
	"bufio"
	"os"
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
	// Parse any labels
	match, err := Label(inp, pointer)
	if err != nil {
		panic(err)
	}
	if match {
		return []byte{}
	}

	// Parse any instructions
	if match, _, result := Either(inp, []Parser{LDA, LDX, LDY}); match {
		pointer++
		return result
	}
	return []byte{}
}
