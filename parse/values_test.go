package parse_test

import (
	"testing"

	"github.com/snarechops/assembler/parse"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	match, result := parse.A("$1234")
	require.True(t, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.A("$65535")
	require.True(t, match)
	require.Equal(t, []byte{0xff, 0xff}, result)

	match, result = parse.A("$0xff12")
	require.True(t, match)
	require.Equal(t, []byte{0x12, 0xff}, result)

	match, result = parse.A("1234")
	require.False(t, match)
}

func TestI(t *testing.T) {
	match, result := parse.I("1234")
	require.True(t, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.I("'c'")
	require.True(t, match)
	require.Equal(t, []byte{0x63}, result)

	match, result = parse.I("0xff")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)
}

func TestAX(t *testing.T) {
	match, result := parse.AX("$1234,x")
	require.True(t, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AX("$1234,X")
	require.True(t, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AX("$0xffd1,x")
	require.True(t, match)
	require.Equal(t, []byte{0xd1, 0xff}, result)

	match, result = parse.AX("$'c',x")
	require.False(t, match)

	match, result = parse.AX("1234,x")
	require.False(t, match)

	match, result = parse.AX("$1234,y")
	require.False(t, match)
}

func TestAY(t *testing.T) {
	match, result := parse.AY("$1234,y")
	require.True(t, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AY("$1234,Y")
	require.True(t, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = parse.AY("$0xffd1,y")
	require.True(t, match)
	require.Equal(t, []byte{0xd1, 0xff}, result)

	match, result = parse.AY("$'c',y")
	require.False(t, match)

	match, result = parse.AY("1234,y")
	require.False(t, match)

	match, result = parse.AY("$1234,x")
	require.False(t, match)
}

func TestZP(t *testing.T) {
	match, result := parse.ZP("$12")
	require.True(t, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.ZP("$0xff")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZP("$255")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZP("$256")
	require.False(t, match)

	match, result = parse.ZP("26")
	require.False(t, match)

	match, result = parse.ZP("$'c'")
	require.False(t, match)

	match, result = parse.ZP("")
	require.False(t, match)
}

func TestZPX(t *testing.T) {
	match, result := parse.ZPX("$12,x")
	require.True(t, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.ZPX("$0xff,x")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPX("$255,x")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPX("$256,x")
	require.False(t, match)

	match, result = parse.ZPX("255,x")
	require.False(t, match)

	match, result = parse.ZPX("$'c',x")
	require.False(t, match)

	match, result = parse.ZPX("")
	require.False(t, match)
}

func TestZPY(t *testing.T) {
	match, result := parse.ZPY("$12,y")
	require.True(t, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.ZPY("$0xff,y")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPY("$255,y")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPY("$256,y")
	require.False(t, match)

	match, result = parse.ZPY("255,y")
	require.False(t, match)

	match, result = parse.ZPY("$'c',y")
	require.False(t, match)

	match, result = parse.ZPY("")
	require.False(t, match)
}

func TestZPIX(t *testing.T) {
	match, result := parse.ZPIX("($12,x)")
	require.True(t, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = parse.ZPIX("($0xff,x)")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPIX("($255,x)")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPIX("$255,x")
	require.False(t, match)

	match, result = parse.ZPIX("($256,x)")
	require.False(t, match)

	match, result = parse.ZPIX("(255,x)")
	require.False(t, match)
}

func TestZPIY(t *testing.T) {
	match, result := parse.ZPIY("($12),y")
	require.True(t, match)
	require.Equal(t, []byte{0xc}, result)

	match, result = parse.ZPIY("($0xff),y")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPIY("($255),y")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.ZPIY("($256),y")
	require.False(t, match)

	match, result = parse.ZPIY("$255,y")
	require.False(t, match)

	match, result = parse.ZPIY("(255),y")
	require.False(t, match)
}
