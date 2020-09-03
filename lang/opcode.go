package lang

// Opcode returns the opcode for the provided
// instruction and address mode
func Opcode(inst string, mode string) []byte {
	return []byte{opcodes[inst][mode]}
}
