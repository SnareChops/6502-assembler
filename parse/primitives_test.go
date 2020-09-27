package parse_test

import (
	"testing"

	"github.com/snarechops/6502-assembler/parse"
	"github.com/stretchr/testify/require"
)

func TestUint8(t *testing.T) {
	match, result := parse.Uint8("6")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0x06}, result)

	match, result = parse.Uint8("255")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.Uint8("256")
	require.Empty(t, match)

	match, result = parse.Uint8("65535")
	require.Empty(t, match)

	match, result = parse.Uint8("x")
	require.Empty(t, match)

}

func TestUint16(t *testing.T) {
	match, result := parse.Uint16("6")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0x06, 0x00}, result)

	match, result = parse.Uint16("255")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0xff, 0x00}, result)

	match, result = parse.Uint16("65535")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0xff, 0xff}, result)

	match, result = parse.Uint16("127236748")
	require.Empty(t, match)

	match, result = parse.Uint16("h")
	require.Empty(t, match)
}

func TestChar(t *testing.T) {
	match, result := parse.Char("'c'")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0x63}, result)

	match, result = parse.Char("'C'")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0x43}, result)

	match, result = parse.Char("' '")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0x20}, result)

	match, result = parse.Char("'''")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0x27}, result)

	match, result = parse.Char("'9'")
	require.Equal(t, `true`, match)
	require.Equal(t, []byte{0x39}, result)

	match, result = parse.Char("'")
	require.Empty(t, match)

	match, result = parse.Char("")
	require.Empty(t, match)

	match, result = parse.Char("'😀'")
	require.Empty(t, match)
}
