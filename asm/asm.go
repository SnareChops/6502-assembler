package asm

import (
	"io"
	"os"

	"github.com/snarechops/6502-assembler/parse"
)

// Assemble assembles the provided asm code
// Returns []byte of binary assembled data
func Assemble(reader io.Reader) ([]byte, error) {
	// Parse ASM
	return parse.Parse(reader)
}

// File assembles the provided asm file
// Returns []byte of binary assembled data
func File(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Assemble(file)
}
