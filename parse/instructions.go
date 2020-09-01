package parse

// LDA parses an LDA instruction
func LDA(inp string) (bool, []byte) {
	return instruction(inp, LDA, `(?i)^LDA\s+(.+)`, ZP, ZPX, AX, AY, A, I, ZPIX, ZPIY)
}

// LDX parses an LDX instruction
func LDX(inp string) (bool, []byte) {
	return instruction(inp, LDX, `(?i)^LDX\s+(.*)`, ZP, ZPY, AY, A, I)
}

// LDY parses an LDY instruction
func LDY(inp string) (bool, []byte) {
	return instruction(inp, LDY, `(?i)^LDY\s+(.*)`, ZP, ZPX, AX, A, I)
}

// STA parses an STA instruction
func STA(inp string) (bool, []byte) {
	return instruction(inp, STA, `(?i)^STA\s+(.*)`, ZP, ZPX, AX, AY, A, ZPIX, ZPIY)
}

// STX parses an STX instruction
func STX(inp string) (bool, []byte) {
	return instruction(inp, STX, `(?i)^STX\s+(.*)`, ZP, ZPY, A)
}

// STY parses an STY instruction
func STY(inp string) (bool, []byte) {
	return instruction(inp, STY, `(?i)^STY\s+(.*)`, ZP, ZPX, A)
}

func instruction(inp string, inst Parser, regex string, parsers ...Parser) (bool, []byte) {
	if match := Submatch(inp, regex); match != nil {
		if valid, mode, value := Either(match[1], parsers); valid {
			return true, append(Opcode(inst, mode), value...)
		}
	}
	return false, nil
}
