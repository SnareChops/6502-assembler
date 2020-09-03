package parse

import (
	"github.com/snarechops/assembler/lang"
)

// LDA parses an LDA instruction
var LDA Parser = inst(lang.LDA, ZP, ZPX, AX, AY, A, I, ZPIX, ZPIY)

// LDX parses an LDX instruction
var LDX Parser = inst(lang.LDX, ZP, ZPY, AY, A, I)

// LDY parses an LDY instruction
var LDY Parser = inst(lang.LDY, ZP, ZPX, AX, A, I)

// STA parses an STA instruction
var STA Parser = inst(lang.STA, ZP, ZPX, AX, AY, A, ZPIX, ZPIY)

// STX parses an STX instruction
var STX Parser = inst(lang.STX, ZP, ZPY, A)

// STY parses an STY instruction
var STY Parser = inst(lang.STY, ZP, ZPX, A)

func inst(acronym string, parsers ...Parser) Parser {
	matcher := Matcher("(?i)^" + acronym + "\\s+([\\w$,()]*)(?:\\s*//)*")
	return func(inp string) (string, []byte) {
		if match := matcher(inp); match != nil {
			if mode, value := Either(match[1], parsers...); mode != "" {
				return acronym, append(lang.Opcode(acronym, mode), value...)
			}
		}
		return "", nil
	}
}
