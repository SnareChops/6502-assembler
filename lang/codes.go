package lang

var opcodes = map[string]map[string]byte{
	LDA: map[string]byte{
		A:    0xad,
		AX:   0xbd,
		AY:   0xb9,
		I:    0xa9,
		ZP:   0xa5,
		ZPIX: 0xa1,
		ZPX:  0xb5,
		ZPIY: 0xb1,
	},
	LDX: map[string]byte{
		A:   0xae,
		AY:  0xbe,
		I:   0xa2,
		ZP:  0xa6,
		ZPY: 0xb6,
	},
	LDY: map[string]byte{
		A:   0xac,
		AX:  0xbc,
		I:   0xa0,
		ZP:  0xa4,
		ZPX: 0xb4,
	},
	STA: map[string]byte{
		A:    0x8d,
		AX:   0x9d,
		AY:   0x99,
		ZP:   0x85,
		ZPIX: 0x81,
		ZPX:  0x95,
		ZPIY: 0x91,
	},
	STX: map[string]byte{
		A:   0x8e,
		ZP:  0x86,
		ZPY: 0x96,
	},
	STY: map[string]byte{
		A:   0x8c,
		ZP:  0x84,
		ZPX: 0x94,
	},
	ADC: map[string]byte{
		A:    0x6d,
		AX:   0x7d,
		AY:   0x79,
		I:    0x69,
		ZP:   0x65,
		ZPIX: 0x61,
		ZPX:  0x75,
		ZPIY: 0x71,
	},
	SBC: map[string]byte{
		A:    0xed,
		AX:   0xfd,
		AY:   0xf9,
		I:    0xe9,
		ZP:   0xe5,
		ZPIX: 0xe1,
		ZPX:  0xf5,
		ZPIY: 0xf1,
	},
	INC: map[string]byte{
		A:   0xee,
		AX:  0xfe,
		ZP:  0xe6,
		ZPX: 0xf6,
	},
	INX: map[string]byte{
		IMP: 0xe8,
	},
	INY: map[string]byte{
		IMP: 0xc8,
	},
	DEC: map[string]byte{
		A:   0xce,
		AX:  0xde,
		ZP:  0xc6,
		ZPX: 0xd6,
	},
	DEX: map[string]byte{
		IMP: 0xca,
	},
	DEY: map[string]byte{
		IMP: 0x88,
	},
	ASL: map[string]byte{
		A:   0x0e,
		AX:  0x1e,
		IMP: 0x0a,
		ZP:  0x06,
		ZPX: 0x16,
	},
	LSR: map[string]byte{
		A:   0x4e,
		AX:  0x5e,
		IMP: 0x4a,
		ZP:  0x46,
		ZPX: 0x56,
	},
	ROL: map[string]byte{
		A:   0x2e,
		AX:  0x3e,
		IMP: 0x2a,
		ZP:  0x26,
		ZPX: 0x36,
	},
	ROR: map[string]byte{
		A:   0x6e,
		AX:  0x7e,
		IMP: 0x6a,
		ZP:  0x66,
		ZPX: 0x76,
	},
	AND: map[string]byte{
		A:    0x2d,
		AX:   0x3d,
		AY:   0x39,
		I:    0x29,
		ZP:   0x25,
		ZPIX: 0x21,
		ZPX:  0x35,
		ZPIY: 0x31,
	},
	ORA: map[string]byte{
		A:    0x0d,
		AX:   0x1d,
		AY:   0x19,
		I:    0x09,
		ZP:   0x05,
		ZPIX: 0x01,
		ZPX:  0x15,
		ZPIY: 0x11,
	},
	XOR: map[string]byte{
		A:    0x4d,
		AX:   0x5d,
		AY:   0x59,
		I:    0x49,
		ZP:   0x45,
		ZPIX: 0x41,
		ZPX:  0x55,
		ZPIY: 0x51,
	},
	CMP: map[string]byte{
		A:    0xcd,
		AX:   0xdd,
		AY:   0xD9,
		I:    0xc9,
		ZP:   0xc5,
		ZPIX: 0xc1,
		ZPX:  0xd5,
		ZPIY: 0xd1,
	},
	CPX: map[string]byte{
		A:  0xec,
		I:  0xe0,
		ZP: 0xe4,
	},
	CPY: map[string]byte{
		A:  0xcc,
		I:  0xc0,
		ZP: 0xc4,
	},
	BIT: map[string]byte{
		A:  0x2c,
		I:  0x89,
		ZP: 0x24,
	},
	BCC: map[string]byte{
		R: 0x90,
	},
	BCS: map[string]byte{
		R: 0xb0,
	},
	BNE: map[string]byte{
		R: 0xd0,
	},
	BEQ: map[string]byte{
		R: 0xf0,
	},
	BPL: map[string]byte{
		R: 0x10,
	},
	BMI: map[string]byte{
		R: 0x30,
	},
	BVC: map[string]byte{
		R: 0x50,
	},
	BVS: map[string]byte{
		R: 0x70,
	},
	TAX: map[string]byte{
		IMP: 0xaa,
	},
	TXA: map[string]byte{
		IMP: 0x8a,
	},
	TAY: map[string]byte{
		IMP: 0xa8,
	},
	TYA: map[string]byte{
		IMP: 0x98,
	},
	TSX: map[string]byte{
		IMP: 0xba,
	},
	TXS: map[string]byte{
		IMP: 0x9a,
	},
	PHA: map[string]byte{
		IMP: 0x48,
	},
	PLA: map[string]byte{
		IMP: 0x68,
	},
	PHP: map[string]byte{
		IMP: 0x08,
	},
	PLP: map[string]byte{
		IMP: 0x28,
	},
	JMP: map[string]byte{
		A:  0x4c,
		AI: 0x6c,
	},
	JSR: map[string]byte{
		A: 0x20,
	},
	RTS: map[string]byte{
		IMP: 0x60,
	},
	RTI: map[string]byte{
		IMP: 0x40,
	},
	CLC: map[string]byte{
		IMP: 0x18,
	},
	SEC: map[string]byte{
		IMP: 0x38,
	},
	CLD: map[string]byte{
		IMP: 0xd8,
	},
	SED: map[string]byte{
		IMP: 0xf8,
	},
	CLI: map[string]byte{
		IMP: 0x58,
	},
	SEI: map[string]byte{
		IMP: 0x78,
	},
	CLV: map[string]byte{
		IMP: 0xb8,
	},
	BRK: map[string]byte{
		IMP: 0x00,
	},
	NOP: map[string]byte{
		IMP: 0xea,
	},
}
