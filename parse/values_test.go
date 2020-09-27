package parse_test

import (
	"testing"

	"github.com/snarechops/6502-assembler/parse"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	match, result := parse.A("$1234")
	require.Equal(t, `A`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.A("$65535")
	require.Equal(t, `A`, match)
	require.Equal(t, []byte{0xff, 0xff}, result)

	match, result = parse.A("$0xff12")
	require.Equal(t, `A`, match)
	require.Equal(t, []byte{0x12, 0xff}, result)

	match, result = parse.A("1234")
	require.Empty(t, match)
}

func TestI(t *testing.T) {
	match, result := parse.I("12")
	require.Equal(t, `I`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.I("'c'")
	require.Equal(t, `I`, match)
	require.Equal(t, []byte{0x63}, result)

	match, result = parse.I("0xff")
	require.Equal(t, `I`, match)
	require.Equal(t, []byte{0xff}, result)
}

func TestAX(t *testing.T) {
	match, result := parse.AX("$1234,x")
	require.Equal(t, `AX`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AX("$1234,X")
	require.Equal(t, `AX`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AX("$0xffd1,x")
	require.Equal(t, `AX`, match)
	require.Equal(t, []byte{0xd1, 0xff}, result)

	match, result = parse.AX("$'c',x")
	require.Empty(t, match)

	match, result = parse.AX("1234,x")
	require.Empty(t, match)

	match, result = parse.AX("$1234,y")
	require.Empty(t, match)
}

func TestAY(t *testing.T) {
	match, result := parse.AY("$1234,y")
	require.Equal(t, `AY`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AY("$1234,Y")
	require.Equal(t, `AY`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AY("$0xffd1,y")
	require.Equal(t, `AY`, match)
	require.Equal(t, []byte{0xd1, 0xff}, result)

	match, result = parse.AY("$'c',y")
	require.Empty(t, match)

	match, result = parse.AY("1234,y")
	require.Empty(t, match)

	match, result = parse.AY("$1234,x")
	require.Empty(t, match)
}

func TestZP(t *testing.T) {
	match, result := parse.ZP("$12")
	require.Equal(t, `ZP`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.ZP("$0xff")
	require.Equal(t, `ZP`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZP("$255")
	require.Equal(t, `ZP`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZP("$256")
	require.Empty(t, match)

	match, result = parse.ZP("26")
	require.Empty(t, match)

	match, result = parse.ZP("$'c'")
	require.Empty(t, match)

	match, result = parse.ZP("")
	require.Empty(t, match)
}

func TestZPX(t *testing.T) {
	match, result := parse.ZPX("$12,x")
	require.Equal(t, `ZPX`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.ZPX("$0xff,x")
	require.Equal(t, `ZPX`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPX("$255,x")
	require.Equal(t, `ZPX`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPX("$256,x")
	require.Empty(t, match)

	match, result = parse.ZPX("255,x")
	require.Empty(t, match)

	match, result = parse.ZPX("$'c',x")
	require.Empty(t, match)

	match, result = parse.ZPX("")
	require.Empty(t, match)
}

func TestZPY(t *testing.T) {
	match, result := parse.ZPY("$12,y")
	require.Equal(t, `ZPY`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.ZPY("$0xff,y")
	require.Equal(t, `ZPY`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPY("$255,y")
	require.Equal(t, `ZPY`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPY("$256,y")
	require.Empty(t, match)

	match, result = parse.ZPY("255,y")
	require.Empty(t, match)

	match, result = parse.ZPY("$'c',y")
	require.Empty(t, match)

	match, result = parse.ZPY("")
	require.Empty(t, match)
}

func TestZPIX(t *testing.T) {
	match, result := parse.ZPIX("($12,x)")
	require.Equal(t, `ZPIX`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.ZPIX("($0xff,x)")
	require.Equal(t, `ZPIX`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPIX("($255,x)")
	require.Equal(t, `ZPIX`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPIX("$255,x")
	require.Empty(t, match)

	match, result = parse.ZPIX("($256,x)")
	require.Empty(t, match)

	match, result = parse.ZPIX("(255,x)")
	require.Empty(t, match)
}

func TestZPIY(t *testing.T) {
	match, result := parse.ZPIY("($12),y")
	require.Equal(t, `ZPIY`, match)
	require.Equal(t, []byte{0xc}, result)

	match, result = parse.ZPIY("($0xff),y")
	require.Equal(t, `ZPIY`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPIY("($255),y")
	require.Equal(t, `ZPIY`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPIY("($256),y")
	require.Empty(t, match)

	match, result = parse.ZPIY("$255,y")
	require.Empty(t, match)

	match, result = parse.ZPIY("(255),y")
	require.Empty(t, match)
}

func TestR(t *testing.T) {
	match, result := parse.R("$12")
	require.Equal(t, `R`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.R("$0x40")
	require.Equal(t, `R`, match)
	require.Equal(t, []byte{0x40}, result)

	match, result = parse.R("$127")
	require.Equal(t, `R`, match)
	require.Equal(t, []byte{0x7f}, result)

	match, result = parse.R("$-128")
	require.Equal(t, `R`, match)
	require.Equal(t, []byte{0x80}, result)

	match, result = parse.R("$128")
	require.Empty(t, match)

	match, result = parse.R("26")
	require.Empty(t, match)

	match, result = parse.R("$'c'")
	require.Empty(t, match)

	match, result = parse.R("")
	require.Empty(t, match)
}

func TestAI(t *testing.T) {
	match, result := parse.AI("($1234)")
	require.Equal(t, `AI`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AI("($65535)")
	require.Equal(t, `AI`, match)
	require.Equal(t, []byte{0xff, 0xff}, result)

	match, result = parse.AI("($0xff12)")
	require.Equal(t, `AI`, match)
	require.Equal(t, []byte{0x12, 0xff}, result)

	match, result = parse.AI("(1234)")
	require.Empty(t, match)
}
