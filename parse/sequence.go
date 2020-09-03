package parse

// SequenceParser the type of a sequence parser
type SequenceParser = func(string, []Parser) (string, []byte)

// Either runs a sequence of parsers returning the result of the
// first matching parser, and a refence to the parser matched
func Either(inp string, parsers ...Parser) (string, []byte) {
	for _, parser := range parsers {
		match, value := parser(inp)
		if match != "" {
			return match, value
		}
	}
	return "", nil
}
