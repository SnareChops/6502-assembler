package parse

import (
	"github.com/snarechops/6502-assembler/lang"
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

// ADC parses an ADC instruction
var ADC Parser = inst(lang.ADC, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)

// SBC parses an SBC instruction
var SBC Parser = inst(lang.SBC, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)

// INC parses an INC instruction
var INC Parser = inst(lang.INC, ZP, ZPX, AX, A)

// INX parses an INX instruction
var INX Parser = implied(lang.INX)

// INY parses an INY instruction
var INY Parser = implied(lang.INY)

// DEC parses a DEC instruction
var DEC Parser = inst(lang.DEC, ZP, ZPX, AX, A)

// DEX parses a DEX instructiom
var DEX Parser = implied(lang.DEX)

// DEY parses a DEY instruction
var DEY Parser = implied(lang.DEY)

// ASL parses an ASL instruction
var ASL Parser = instOrImp(lang.ASL, ZP, ZPX, AX, A)

// LSR parses an LSR instruction
var LSR Parser = instOrImp(lang.LSR, ZP, ZPX, AX, A)

// ROL parses an ROL instruction
var ROL Parser = instOrImp(lang.ROL, ZP, ZPX, AX, A)

// ROR parses an ROR instruction
var ROR Parser = instOrImp(lang.ROR, ZP, ZPX, AX, A)

// AND parses an AND instruction
var AND Parser = inst(lang.AND, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)

// ORA parses an ORA instruction
var ORA Parser = inst(lang.ORA, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)

// XOR parses an XOR instruction
var XOR Parser = inst(lang.XOR, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)

// CMP parses a CMP instruction
var CMP Parser = inst(lang.CMP, ZP, ZPX, ZPIX, ZPIY, AX, AY, A, I)

// CPX parses a CPX instruction
var CPX Parser = inst(lang.CPX, ZP, A, I)

// CPY parses a CPY instruction
var CPY Parser = inst(lang.CPY, ZP, A, I)

// BIT parses a BIT instruction
var BIT Parser = inst(lang.BIT, ZP, A, I)

// BCC parses a BCC instruction
var BCC Parser = inst(lang.BCC, R)

// BCS parses a BCS instruction
var BCS Parser = inst(lang.BCS, R)

// BNE parses a BNE instruction
var BNE Parser = inst(lang.BNE, R)

// BEQ parses a BEQ instruction
var BEQ Parser = inst(lang.BEQ, R)

// BPL parses a BPL instruction
var BPL Parser = inst(lang.BPL, R)

// BMI parses a BMI instruction
var BMI Parser = inst(lang.BMI, R)

// BVC parses a BVC instruction
var BVC Parser = inst(lang.BVC, R)

// BVS parses a BVS instruction
var BVS Parser = inst(lang.BVS, R)

// TAX parses a TAX instruction
var TAX Parser = implied(lang.TAX)

// TXA parses a TXA instruction
var TXA Parser = implied(lang.TXA)

// TAY parses a TAY instruction
var TAY Parser = implied(lang.TAY)

// TYA parses a TYA instruction
var TYA Parser = implied(lang.TYA)

// TSX parses a TSX instruction
var TSX Parser = implied(lang.TSX)

// TXS parses a TXS instruction
var TXS Parser = implied(lang.TXS)

// PHA parses a PHA instruction
var PHA Parser = implied(lang.PHA)

// PLA parses a PLA instruction
var PLA Parser = implied(lang.PLA)

// PHP parses a PHP instruction
var PHP Parser = implied(lang.PHP)

// PLP parses a PLP instruction
var PLP Parser = implied(lang.PLP)

// JMP parses a JMP instruction
var JMP Parser = inst(lang.JMP, AI, A)

// JSR parses a JSR instruction
var JSR Parser = inst(lang.JSR, A)

// RTS parses an RTS instruction
var RTS Parser = implied(lang.RTS)

// RTI parses an RTI instruction
var RTI Parser = implied(lang.RTI)

// CLC parses a CLC instruction
var CLC Parser = implied(lang.CLC)

// SEC parses an SEC instruction
var SEC Parser = implied(lang.SEC)

// CLD parses a CLD instruction
var CLD Parser = implied(lang.CLD)

// SED parses an SED instruction
var SED Parser = implied(lang.SED)

// CLI parses a CLI instruction
var CLI Parser = implied(lang.CLI)

// SEI parses an SEI instruction
var SEI Parser = implied(lang.SEI)

// CLV parses a CLV instruction
var CLV Parser = implied(lang.CLV)

// BRK parses a BRK instruction
var BRK Parser = implied(lang.BRK)

// NOP parses a NOP instruction
var NOP Parser = implied(lang.NOP)

// Instructions List of all instruction parsers
var Instructions = []Parser{LDA, LDX, LDY, STA, STX, STY, ADC, SBC, INC, INX, INY, DEC, DEX, DEY, ASL, LSR, ROL, ROR, AND, ORA, XOR, CMP, CPX, CPY, BIT, BCC, BCS, BNE, BEQ, BPL, BMI, BVC, BVS, TAX, TXA, TAY, TYA, TSX, TXS, PHA, PLA, PHP, PLP, JMP, JSR, RTS, RTI, CLC, SEC, CLD, SED, CLI, SEI, CLV, BRK, NOP}

func inst(acronym string, parsers ...Parser) Parser {
	matcher := Matcher("(?i)^" + acronym + "\\s+([\\w$,()]*)(?:\\s*//.*)*$")
	return func(inp string) (string, []byte) {
		if match := matcher(inp); match != nil {
			match[1] = ReplaceFromLabel(match[1])
			if mode, value := Either(match[1], parsers...); mode != "" {
				return acronym, append(lang.Opcode(acronym, mode), value...)
			}
		}
		return "", nil
	}
}

func implied(acronym string) Parser {
	matcher := Matcher("(?i)^" + acronym + "(?:\\s*//.*)*$")
	return func(inp string) (string, []byte) {
		if match := matcher(inp); match != nil {
			return acronym, lang.Opcode(acronym, lang.IMP)
		}
		return "", nil
	}
}

func instOrImp(acronym string, parsers ...Parser) Parser {
	in := inst(acronym, parsers...)
	imp := implied(acronym)
	return func(inp string) (string, []byte) {
		return Either(inp, imp, in)
	}
}
