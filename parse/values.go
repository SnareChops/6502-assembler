package parse

// Parser represents a parser function
type Parser = func(string) (bool, []byte)

// A parses an absolute memory address
func A(inp string) (bool, []byte) {
	if match := Submatch(inp, `^\$(.+)`); match != nil {
		if valid, value := Uint16(match[1]); valid {
			return true, value
		}
	}
	return false, nil
}

// I parses an immediate literal value
func I(inp string) (bool, []byte) {
	valid, _, value := Either(inp, []Parser{Char, Uint8, Uint16})
	return valid, value
}

// AX parses an absolute with x value
// ex: $1234,x
func AX(inp string) (bool, []byte) {
	if match := Submatch(inp, `(?i)^\$(.+),x$`); match != nil {
		if valid, value := Uint16(match[1]); valid {
			return true, value
		}
	}
	return false, nil
}

// AY parses an absolute with y value
// ex: $1234,y
func AY(inp string) (bool, []byte) {
	if match := Submatch(inp, `(?i)^\$(.+),y$`); match != nil {
		if valid, value := Uint16(match[1]); valid {
			return true, value
		}
	}
	return false, nil
}

// ZP parses a zero page address
func ZP(inp string) (bool, []byte) {
	if match := Submatch(inp, `^\$(.+)$`); match != nil {
		if valid, value := Uint8(match[1]); valid {
			return true, value
		}
	}
	return false, nil
}

// ZPX parses a zero paged with x address
func ZPX(inp string) (bool, []byte) {
	if match := Submatch(inp, `^(.+),x$`); match != nil {
		if valid, value := ZP(match[1]); valid {
			return true, value
		}
	}
	return false, nil
}

// ZPY parses a zero paged with y address
func ZPY(inp string) (bool, []byte) {
	if match := Submatch(inp, `^(.+),y$`); match != nil {
		if valid, value := ZP(match[1]); valid {
			return true, value
		}
	}
	return false, nil
}

// ZPIX parses a zero page indirect indexed with x address
func ZPIX(inp string) (bool, []byte) {
	if match := Submatch(inp, `^\((.+),x\)$`); match != nil {
		if valid, value := ZP(match[1]); valid {
			return true, value
		}
	}
	return false, nil
}

// ZPIY parses a zero page indirect indexed with y address
func ZPIY(inp string) (bool, []byte) {
	if match := Submatch(inp, `^\((.+)\),y$`); match != nil {
		if valid, value := ZP(match[1]); valid {
			return true, value
		}
	}
	return false, nil
}
