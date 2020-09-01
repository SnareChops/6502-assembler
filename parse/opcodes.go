package parse

import (
	"reflect"
	"runtime"
	"strings"
)

var opcodes = map[string]map[string]byte{
	"LDA": map[string]byte{
		"A":    0xad,
		"AX":   0xbd,
		"AY":   0xb9,
		"I":    0xa9,
		"ZP":   0xa5,
		"ZPIX": 0xa1,
		"ZPX":  0xb5,
		"ZPIY": 0xb1,
	},
	"LDX": map[string]byte{
		"A":   0xae,
		"AY":  0xbe,
		"I":   0xa2,
		"ZP":  0xa6,
		"ZPY": 0xb6,
	},
	"LDY": map[string]byte{
		"A":   0xac,
		"AX":  0xbc,
		"I":   0xa0,
		"ZP":  0xa4,
		"ZPX": 0xb4,
	},
	"STA": map[string]byte{
		"A":    0x8d,
		"AX":   0x9d,
		"AY":   0x99,
		"ZP":   0x85,
		"ZPIX": 0x81,
		"ZPX":  0x95,
		"ZPIY": 0x91,
	},
	"STX": map[string]byte{
		"A":   0x8e,
		"ZP":  0x86,
		"ZPY": 0x96,
	},
	"STY": map[string]byte{
		"A":   0x8c,
		"ZP":  0x84,
		"ZPX": 0x94,
	},
}

// Opcode returns the matching opcode for the
// instruction and memory address mode
func Opcode(inst Parser, mode Parser) []byte {
	return []byte{opcodes[funcName(inst)][funcName(mode)]}
}

func funcName(f Parser) string {
	long := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	parts := strings.Split(long, ".")
	name := parts[len(parts)-1]
	return name
}
