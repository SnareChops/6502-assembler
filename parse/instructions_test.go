package parse_test

import (
	"testing"

	"github.com/snarechops/6502-assembler/parse"
	"github.com/stretchr/testify/require"
)

func TestLDA(t *testing.T) {
	match, result := parse.LDA("LDA $12")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xa5, 0x0c}, result)

	match, result = parse.LDA("lda $12")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xa5, 0x0c}, result)

	match, result = parse.LDA("LDA $12,x")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xb5, 0x0c}, result)

	match, result = parse.LDA("LDA ($12,x)")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xa1, 0x0c}, result)

	match, result = parse.LDA("LDA ($12),y")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xb1, 0x0c}, result)

	match, result = parse.LDA("LDA $0xff12")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xad, 0x12, 0xff}, result)

	match, result = parse.LDA("LDA $0xff12,x")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xbd, 0x12, 0xff}, result)

	match, result = parse.LDA("LDA $0xff12,y")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xb9, 0x12, 0xff}, result)

	match, result = parse.LDA("LDA 0x12")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xa9, 0x12}, result)
}

func TestLDX(t *testing.T) {
	match, result := parse.LDX("LDX $0xff12")
	require.Equal(t, `LDX`, match)
	require.Equal(t, []byte{0xae, 0x12, 0xff}, result)

	match, result = parse.LDX("LDX $0xff12,y")
	require.Equal(t, `LDX`, match)
	require.Equal(t, []byte{0xbe, 0x12, 0xff}, result)

	match, result = parse.LDX("LDX 0x12")
	require.Equal(t, `LDX`, match)
	require.Equal(t, []byte{0xa2, 0x12}, result)

	match, result = parse.LDX("LDX $12")
	require.Equal(t, `LDX`, match)
	require.Equal(t, []byte{0xa6, 0x0c}, result)

	match, result = parse.LDX("LDX $12,y")
	require.Equal(t, `LDX`, match)
	require.Equal(t, []byte{0xb6, 0x0c}, result)
}

func TestLDY(t *testing.T) {
	match, result := parse.LDY("LDY $0xff12")
	require.Equal(t, `LDY`, match)
	require.Equal(t, []byte{0xac, 0x12, 0xff}, result)

	match, result = parse.LDY("LDY $0xff12,x")
	require.Equal(t, `LDY`, match)
	require.Equal(t, []byte{0xbc, 0x12, 0xff}, result)

	match, result = parse.LDY("LDY 0x12")
	require.Equal(t, `LDY`, match)
	require.Equal(t, []byte{0xa0, 0x12}, result)

	match, result = parse.LDY("LDY $12")
	require.Equal(t, `LDY`, match)
	require.Equal(t, []byte{0xa4, 0x0c}, result)

	match, result = parse.LDY("LDY $12,x")
	require.Equal(t, `LDY`, match)
	require.Equal(t, []byte{0xb4, 0x0c}, result)
}

func TestSTA(t *testing.T) {
	match, result := parse.STA("STA $0xff12")
	require.Equal(t, `STA`, match)
	require.Equal(t, []byte{0x8d, 0x12, 0xff}, result)

	match, result = parse.STA("STA $0xff12,x")
	require.Equal(t, `STA`, match)
	require.Equal(t, []byte{0x9d, 0x12, 0xff}, result)

	match, result = parse.STA("STA $0xff12,y")
	require.Equal(t, `STA`, match)
	require.Equal(t, []byte{0x99, 0x12, 0xff}, result)

	match, result = parse.STA("STA $12")
	require.Equal(t, `STA`, match)
	require.Equal(t, []byte{0x85, 0x0c}, result)

	match, result = parse.STA("STA ($12,x)")
	require.Equal(t, `STA`, match)
	require.Equal(t, []byte{0x81, 0x0c}, result)

	match, result = parse.STA("STA $12,x")
	require.Equal(t, `STA`, match)
	require.Equal(t, []byte{0x95, 0x0c}, result)

	match, result = parse.STA("STA ($12),y")
	require.Equal(t, `STA`, match)
	require.Equal(t, []byte{0x91, 0x0c}, result)
}

func TestSTX(t *testing.T) {
	match, result := parse.STX("STX $0xff12")
	require.Equal(t, `STX`, match)
	require.Equal(t, []byte{0x8e, 0x12, 0xff}, result)

	match, result = parse.STX("STX $12")
	require.Equal(t, `STX`, match)
	require.Equal(t, []byte{0x86, 0x0c}, result)

	match, result = parse.STX("STX $12,y")
	require.Equal(t, `STX`, match)
	require.Equal(t, []byte{0x96, 0x0c}, result)
}

func TestSTY(t *testing.T) {
	match, result := parse.STY("STY $0xff12")
	require.Equal(t, `STY`, match)
	require.Equal(t, []byte{0x8c, 0x12, 0xff}, result)

	match, result = parse.STY("STY $12")
	require.Equal(t, `STY`, match)
	require.Equal(t, []byte{0x84, 0x0c}, result)

	match, result = parse.STY("STY $12,x")
	require.Equal(t, `STY`, match)
	require.Equal(t, []byte{0x94, 0x0c}, result)
}

func TestADC(t *testing.T) {
	match, result := parse.ADC("ADC $0xff12")
	require.Equal(t, `ADC`, match)
	require.Equal(t, []byte{0x6d, 0x12, 0xff}, result)

	match, result = parse.ADC("ADC $0xff12,x")
	require.Equal(t, `ADC`, match)
	require.Equal(t, []byte{0x7d, 0x12, 0xff}, result)

	match, result = parse.ADC("ADC $0xff12,y")
	require.Equal(t, `ADC`, match)
	require.Equal(t, []byte{0x79, 0x12, 0xff}, result)

	match, result = parse.ADC("ADC 0x12")
	require.Equal(t, `ADC`, match)
	require.Equal(t, []byte{0x69, 0x12}, result)

	match, result = parse.ADC("ADC $12")
	require.Equal(t, `ADC`, match)
	require.Equal(t, []byte{0x65, 0x0c}, result)

	match, result = parse.ADC("ADC ($12,x)")
	require.Equal(t, `ADC`, match)
	require.Equal(t, []byte{0x61, 0x0c}, result)

	match, result = parse.ADC("ADC $12,x")
	require.Equal(t, `ADC`, match)
	require.Equal(t, []byte{0x75, 0x0c}, result)

	match, result = parse.ADC("ADC ($12),y")
	require.Equal(t, `ADC`, match)
	require.Equal(t, []byte{0x71, 0x0c}, result)
}

func TestSBC(t *testing.T) {
	match, result := parse.SBC("SBC $0xff12")
	require.Equal(t, `SBC`, match)
	require.Equal(t, []byte{0xed, 0x12, 0xff}, result)

	match, result = parse.SBC("SBC $0xff12,x")
	require.Equal(t, `SBC`, match)
	require.Equal(t, []byte{0xfd, 0x12, 0xff}, result)

	match, result = parse.SBC("SBC $0xff12,y")
	require.Equal(t, `SBC`, match)
	require.Equal(t, []byte{0xf9, 0x12, 0xff}, result)

	match, result = parse.SBC("SBC 0x12")
	require.Equal(t, `SBC`, match)
	require.Equal(t, []byte{0xe9, 0x12}, result)

	match, result = parse.SBC("SBC $12")
	require.Equal(t, `SBC`, match)
	require.Equal(t, []byte{0xe5, 0x0c}, result)

	match, result = parse.SBC("SBC ($12,x)")
	require.Equal(t, `SBC`, match)
	require.Equal(t, []byte{0xe1, 0x0c}, result)

	match, result = parse.SBC("SBC $12,x")
	require.Equal(t, `SBC`, match)
	require.Equal(t, []byte{0xf5, 0x0c}, result)

	match, result = parse.SBC("SBC ($12),y")
	require.Equal(t, `SBC`, match)
	require.Equal(t, []byte{0xf1, 0x0c}, result)
}

func TestINC(t *testing.T) {
	match, result := parse.INC("INC $0xff12")
	require.Equal(t, `INC`, match)
	require.Equal(t, []byte{0xee, 0x12, 0xff}, result)

	match, result = parse.INC("INC $0xff12,x")
	require.Equal(t, `INC`, match)
	require.Equal(t, []byte{0xfe, 0x12, 0xff}, result)

	match, result = parse.INC("INC $12")
	require.Equal(t, `INC`, match)
	require.Equal(t, []byte{0xe6, 0x0c}, result)

	match, result = parse.INC("INC $12,x")
	require.Equal(t, `INC`, match)
	require.Equal(t, []byte{0xf6, 0x0c}, result)
}

func TestINX(t *testing.T) {
	match, result := parse.INX("INX")
	require.Equal(t, `INX`, match)
	require.Equal(t, []byte{0xe8}, result)
}

func TestINY(t *testing.T) {
	match, result := parse.INY("INY")
	require.Equal(t, `INY`, match)
	require.Equal(t, []byte{0xc8}, result)
}

func TestDEC(t *testing.T) {
	match, result := parse.DEC("DEC $0xff12")
	require.Equal(t, `DEC`, match)
	require.Equal(t, []byte{0xce, 0x12, 0xff}, result)

	match, result = parse.DEC("DEC $0xff12,x")
	require.Equal(t, `DEC`, match)
	require.Equal(t, []byte{0xde, 0x12, 0xff}, result)

	match, result = parse.DEC("DEC $12")
	require.Equal(t, `DEC`, match)
	require.Equal(t, []byte{0xc6, 0x0c}, result)

	match, result = parse.DEC("DEC $12,x")
	require.Equal(t, `DEC`, match)
	require.Equal(t, []byte{0xd6, 0x0c}, result)
}

func TestDEX(t *testing.T) {
	match, result := parse.DEX("DEX")
	require.Equal(t, `DEX`, match)
	require.Equal(t, []byte{0xca}, result)
}

func TestDEY(t *testing.T) {
	match, result := parse.DEY("DEY")
	require.Equal(t, `DEY`, match)
	require.Equal(t, []byte{0x88}, result)
}

func TestASL(t *testing.T) {
	match, result := parse.ASL("ASL $0xff12")
	require.Equal(t, `ASL`, match)
	require.Equal(t, []byte{0x0e, 0x12, 0xff}, result)

	match, result = parse.ASL("ASL $0xff12,x")
	require.Equal(t, `ASL`, match)
	require.Equal(t, []byte{0x1e, 0x12, 0xff}, result)

	match, result = parse.ASL("ASL")
	require.Equal(t, `ASL`, match)
	require.Equal(t, []byte{0x0a}, result)

	match, result = parse.ASL("ASL $12")
	require.Equal(t, `ASL`, match)
	require.Equal(t, []byte{0x06, 0x0c}, result)

	match, result = parse.ASL("ASL $12,x")
	require.Equal(t, `ASL`, match)
	require.Equal(t, []byte{0x16, 0x0c}, result)
}

func TestLSR(t *testing.T) {
	match, result := parse.LSR("LSR $0xff12")
	require.Equal(t, `LSR`, match)
	require.Equal(t, []byte{0x4e, 0x12, 0xff}, result)

	match, result = parse.LSR("LSR $0xff12,x")
	require.Equal(t, `LSR`, match)
	require.Equal(t, []byte{0x5e, 0x12, 0xff}, result)

	match, result = parse.LSR("LSR")
	require.Equal(t, `LSR`, match)
	require.Equal(t, []byte{0x4a}, result)

	match, result = parse.LSR("LSR $12")
	require.Equal(t, `LSR`, match)
	require.Equal(t, []byte{0x46, 0x0c}, result)

	match, result = parse.LSR("LSR $12,x")
	require.Equal(t, `LSR`, match)
	require.Equal(t, []byte{0x56, 0x0c}, result)
}

func TestROL(t *testing.T) {
	match, result := parse.ROL("ROL $0xff12")
	require.Equal(t, `ROL`, match)
	require.Equal(t, []byte{0x2e, 0x12, 0xff}, result)

	match, result = parse.ROL("ROL $0xff12,x")
	require.Equal(t, `ROL`, match)
	require.Equal(t, []byte{0x3e, 0x12, 0xff}, result)

	match, result = parse.ROL("ROL")
	require.Equal(t, `ROL`, match)
	require.Equal(t, []byte{0x2a}, result)

	match, result = parse.ROL("ROL $12")
	require.Equal(t, `ROL`, match)
	require.Equal(t, []byte{0x26, 0x0c}, result)

	match, result = parse.ROL("ROL $12,x")
	require.Equal(t, `ROL`, match)
	require.Equal(t, []byte{0x36, 0x0c}, result)
}

func ROR(t *testing.T) {
	match, result := parse.ROR("ROR $0xff12")
	require.Equal(t, `ROR`, match)
	require.Equal(t, []byte{0x6e, 0x12, 0xff}, result)

	match, result = parse.ROR("ROR $0xff12,x")
	require.Equal(t, `ROR`, match)
	require.Equal(t, []byte{0x7e, 0x12, 0xff}, result)

	match, result = parse.ROR("ROR")
	require.Equal(t, `ROR`, match)
	require.Equal(t, []byte{0x6a}, result)

	match, result = parse.ROR("ROR $12")
	require.Equal(t, `ROR`, match)
	require.Equal(t, []byte{0x66, 0x0c}, result)

	match, result = parse.ROR("ROR $12,x")
	require.Equal(t, `ROR`, match)
	require.Equal(t, []byte{0x76, 0x0c}, result)
}

func TestAND(t *testing.T) {
	match, result := parse.AND("AND $0xff12")
	require.Equal(t, `AND`, match)
	require.Equal(t, []byte{0x2d, 0x12, 0xff}, result)

	match, result = parse.AND("AND $0xff12,x")
	require.Equal(t, `AND`, match)
	require.Equal(t, []byte{0x3d, 0x12, 0xff}, result)

	match, result = parse.AND("AND $0xff12,y")
	require.Equal(t, `AND`, match)
	require.Equal(t, []byte{0x39, 0x12, 0xff}, result)

	match, result = parse.AND("AND 0x12")
	require.Equal(t, `AND`, match)
	require.Equal(t, []byte{0x29, 0x12}, result)

	match, result = parse.AND("AND $12")
	require.Equal(t, `AND`, match)
	require.Equal(t, []byte{0x25, 0x0c}, result)

	match, result = parse.AND("AND ($12,x)")
	require.Equal(t, `AND`, match)
	require.Equal(t, []byte{0x21, 0x0c}, result)

	match, result = parse.AND("AND $12,x")
	require.Equal(t, `AND`, match)
	require.Equal(t, []byte{0x35, 0x0c}, result)

	match, result = parse.AND("AND ($12),y")
	require.Equal(t, `AND`, match)
	require.Equal(t, []byte{0x31, 0x0c}, result)
}

func TestORA(t *testing.T) {
	match, result := parse.ORA("ORA $0xff12")
	require.Equal(t, `ORA`, match)
	require.Equal(t, []byte{0x0d, 0x012, 0xff}, result)

	match, result = parse.ORA("ORA $0xff12,x")
	require.Equal(t, `ORA`, match)
	require.Equal(t, []byte{0x1d, 0x12, 0xff}, result)

	match, result = parse.ORA("ORA $0xff12,y")
	require.Equal(t, `ORA`, match)
	require.Equal(t, []byte{0x19, 0x12, 0xff}, result)

	match, result = parse.ORA("ORA 0x12")
	require.Equal(t, `ORA`, match)
	require.Equal(t, []byte{0x09, 0x12}, result)

	match, result = parse.ORA("ORA $12")
	require.Equal(t, `ORA`, match)
	require.Equal(t, []byte{0x05, 0x0c}, result)

	match, result = parse.ORA("ORA ($12,x)")
	require.Equal(t, `ORA`, match)
	require.Equal(t, []byte{0x01, 0x0c}, result)

	match, result = parse.ORA("ORA $12,x")
	require.Equal(t, `ORA`, match)
	require.Equal(t, []byte{0x15, 0x0c}, result)

	match, result = parse.ORA("ORA ($12),y")
	require.Equal(t, `ORA`, match)
	require.Equal(t, []byte{0x11, 0x0c}, result)
}

func TestXOR(t *testing.T) {
	match, result := parse.XOR("XOR $0xff12")
	require.Equal(t, `XOR`, match)
	require.Equal(t, []byte{0x4d, 0x12, 0xff}, result)

	match, result = parse.XOR("XOR $0xff12,x")
	require.Equal(t, `XOR`, match)
	require.Equal(t, []byte{0x5d, 0x12, 0xff}, result)

	match, result = parse.XOR("XOR $0xff12,y")
	require.Equal(t, `XOR`, match)
	require.Equal(t, []byte{0x59, 0x12, 0xff}, result)

	match, result = parse.XOR("XOR 0x12")
	require.Equal(t, `XOR`, match)
	require.Equal(t, []byte{0x49, 0x12}, result)

	match, result = parse.XOR("XOR $12")
	require.Equal(t, `XOR`, match)
	require.Equal(t, []byte{0x45, 0x0c}, result)

	match, result = parse.XOR("XOR ($12,x)")
	require.Equal(t, `XOR`, match)
	require.Equal(t, []byte{0x41, 0x0c}, result)

	match, result = parse.XOR("XOR $12,x")
	require.Equal(t, `XOR`, match)
	require.Equal(t, []byte{0x55, 0x0c}, result)

	match, result = parse.XOR("XOR ($12),y")
	require.Equal(t, `XOR`, match)
	require.Equal(t, []byte{0x51, 0x0c}, result)
}

func TestCMP(t *testing.T) {
	match, result := parse.CMP("CMP $0xff12")
	require.Equal(t, `CMP`, match)
	require.Equal(t, []byte{0xcd, 0x12, 0xff}, result)

	match, result = parse.CMP("CMP $0xff12,x")
	require.Equal(t, `CMP`, match)
	require.Equal(t, []byte{0xdd, 0x12, 0xff}, result)

	match, result = parse.CMP("CMP $0xff12,y")
	require.Equal(t, `CMP`, match)
	require.Equal(t, []byte{0xd9, 0x12, 0xff}, result)

	match, result = parse.CMP("CMP 0x12")
	require.Equal(t, `CMP`, match)
	require.Equal(t, []byte{0xc9, 0x12}, result)

	match, result = parse.CMP("CMP $12")
	require.Equal(t, `CMP`, match)
	require.Equal(t, []byte{0xc5, 0x0c}, result)

	match, result = parse.CMP("CMP ($12,x)")
	require.Equal(t, `CMP`, match)
	require.Equal(t, []byte{0xc1, 0x0c}, result)

	match, result = parse.CMP("CMP $12,x")
	require.Equal(t, `CMP`, match)
	require.Equal(t, []byte{0xd5, 0x0c}, result)

	match, result = parse.CMP("CMP ($12),y")
	require.Equal(t, `CMP`, match)
	require.Equal(t, []byte{0xd1, 0x0c}, result)
}

func TestCPX(t *testing.T) {
	match, result := parse.CPX("CPX $0xff12")
	require.Equal(t, `CPX`, match)
	require.Equal(t, []byte{0xec, 0x12, 0xff}, result)

	match, result = parse.CPX("CPX 0x12")
	require.Equal(t, `CPX`, match)
	require.Equal(t, []byte{0xe0, 0x12}, result)

	match, result = parse.CPX("CPX $12")
	require.Equal(t, `CPX`, match)
	require.Equal(t, []byte{0xe4, 0x0c}, result)
}

func TestCPY(t *testing.T) {
	match, result := parse.CPY("CPY $0xff12")
	require.Equal(t, `CPY`, match)
	require.Equal(t, []byte{0xcc, 0x12, 0xff}, result)

	match, result = parse.CPY("CPY 0x12")
	require.Equal(t, `CPY`, match)
	require.Equal(t, []byte{0xc0, 0x12}, result)

	match, result = parse.CPY("CPY $12")
	require.Equal(t, `CPY`, match)
	require.Equal(t, []byte{0xc4, 0x0c}, result)
}

func TestBIT(t *testing.T) {
	match, result := parse.BIT("BIT $0xff12")
	require.Equal(t, `BIT`, match)
	require.Equal(t, []byte{0x2c, 0x12, 0xff}, result)

	match, result = parse.BIT("BIT 0x12")
	require.Equal(t, `BIT`, match)
	require.Equal(t, []byte{0x89, 0x12}, result)

	match, result = parse.BIT("BIT $12")
	require.Equal(t, `BIT`, match)
	require.Equal(t, []byte{0x24, 0x0c}, result)
}

func TestBCC(t *testing.T) {
	match, result := parse.BCC("BCC $12")
	require.Equal(t, `BCC`, match)
	require.Equal(t, []byte{0x90, 0x0c}, result)
}

func TestBCS(t *testing.T) {
	match, result := parse.BCS("BCS $12")
	require.Equal(t, `BCS`, match)
	require.Equal(t, []byte{0xb0, 0x0c}, result)
}

func TestBNE(t *testing.T) {
	match, result := parse.BNE("BNE $12")
	require.Equal(t, `BNE`, match)
	require.Equal(t, []byte{0xd0, 0x0c}, result)
}

func TestBEQ(t *testing.T) {
	match, result := parse.BEQ("BEQ $12")
	require.Equal(t, `BEQ`, match)
	require.Equal(t, []byte{0xf0, 0x0c}, result)
}

func TestBPL(t *testing.T) {
	match, result := parse.BPL("BPL $12")
	require.Equal(t, `BPL`, match)
	require.Equal(t, []byte{0x10, 0x0c}, result)
}

func TestBMI(t *testing.T) {
	match, result := parse.BMI("BMI $12")
	require.Equal(t, `BMI`, match)
	require.Equal(t, []byte{0x30, 0x0c}, result)
}

func TestBVC(t *testing.T) {
	match, result := parse.BVC("BVC $12")
	require.Equal(t, `BVC`, match)
	require.Equal(t, []byte{0x50, 0x0c}, result)
}

func TestBVS(t *testing.T) {
	match, result := parse.BVS("BVS $12")
	require.Equal(t, `BVS`, match)
	require.Equal(t, []byte{0x70, 0x0c}, result)
}

func TestTAX(t *testing.T) {
	match, result := parse.TAX("TAX")
	require.Equal(t, `TAX`, match)
	require.Equal(t, []byte{0xaa}, result)
}

func TestTXA(t *testing.T) {
	match, result := parse.TXA("TXA")
	require.Equal(t, `TXA`, match)
	require.Equal(t, []byte{0x8a}, result)
}

func TestTAY(t *testing.T) {
	match, result := parse.TAY("TAY")
	require.Equal(t, `TAY`, match)
	require.Equal(t, []byte{0xa8}, result)
}

func TestTYA(t *testing.T) {
	match, result := parse.TYA("TYA")
	require.Equal(t, `TYA`, match)
	require.Equal(t, []byte{0x98}, result)
}

func TestTSX(t *testing.T) {
	match, result := parse.TSX("TSX")
	require.Equal(t, `TSX`, match)
	require.Equal(t, []byte{0xba}, result)
}

func TestTXS(t *testing.T) {
	match, result := parse.TXS("TXS")
	require.Equal(t, `TXS`, match)
	require.Equal(t, []byte{0x9a}, result)
}

func TestPHA(t *testing.T) {
	match, result := parse.PHA("PHA")
	require.Equal(t, `PHA`, match)
	require.Equal(t, []byte{0x48}, result)
}

func TestPLA(t *testing.T) {
	match, result := parse.PLA("PLA")
	require.Equal(t, `PLA`, match)
	require.Equal(t, []byte{0x68}, result)
}

func TestPHP(t *testing.T) {
	match, result := parse.PHP("PHP")
	require.Equal(t, `PHP`, match)
	require.Equal(t, []byte{0x08}, result)
}

func TestPLP(t *testing.T) {
	match, result := parse.PLP("PLP")
	require.Equal(t, `PLP`, match)
	require.Equal(t, []byte{0x28}, result)
}

func TestJMP(t *testing.T) {
	match, result := parse.JMP("JMP $0xff12")
	require.Equal(t, `JMP`, match)
	require.Equal(t, []byte{0x4c, 0x12, 0xff}, result)

	match, result = parse.JMP("JMP ($0xff12)")
	require.Equal(t, `JMP`, match)
	require.Equal(t, []byte{0x6c, 0x12, 0xff}, result)
}

func TestJSR(t *testing.T) {
	match, result := parse.JSR("JSR $0xff12")
	require.Equal(t, `JSR`, match)
	require.Equal(t, []byte{0x20, 0x12, 0xff}, result)
}

func TestRTS(t *testing.T) {
	match, result := parse.RTS("RTS")
	require.Equal(t, `RTS`, match)
	require.Equal(t, []byte{0x60}, result)
}

func TestRTI(t *testing.T) {
	match, result := parse.RTI("RTI")
	require.Equal(t, `RTI`, match)
	require.Equal(t, []byte{0x40}, result)
}

func TestCLC(t *testing.T) {
	match, result := parse.CLC("CLC")
	require.Equal(t, `CLC`, match)
	require.Equal(t, []byte{0x18}, result)
}

func TestSEC(t *testing.T) {
	match, result := parse.SEC("SEC")
	require.Equal(t, `SEC`, match)
	require.Equal(t, []byte{0x38}, result)
}

func TestCLD(t *testing.T) {
	match, result := parse.CLD("CLD")
	require.Equal(t, `CLD`, match)
	require.Equal(t, []byte{0xd8}, result)
}

func TestSED(t *testing.T) {
	match, result := parse.SED("SED")
	require.Equal(t, `SED`, match)
	require.Equal(t, []byte{0xf8}, result)
}

func TestCLI(t *testing.T) {
	match, result := parse.CLI("CLI")
	require.Equal(t, `CLI`, match)
	require.Equal(t, []byte{0x58}, result)
}

func TestSEI(t *testing.T) {
	match, result := parse.SEI("SEI")
	require.Equal(t, `SEI`, match)
	require.Equal(t, []byte{0x78}, result)
}

func TestCLV(t *testing.T) {
	match, result := parse.CLV("CLV")
	require.Equal(t, `CLV`, match)
	require.Equal(t, []byte{0xb8}, result)
}

func TestBRK(t *testing.T) {
	match, result := parse.BRK("BRK")
	require.Equal(t, `BRK`, match)
	require.Equal(t, []byte{0x00}, result)
}

func TestNOP(t *testing.T) {
	match, result := parse.NOP("NOP")
	require.Equal(t, `NOP`, match)
	require.Equal(t, []byte{0xea}, result)
}
