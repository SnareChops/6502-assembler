package parse_test

import (
	"testing"

	"github.com/snarechops/assembler/parse"
	"github.com/stretchr/testify/require"
)

func TestUint8(t *testing.T) {
	match, result := parse.Uint8("6")
	require.True(t, match)
	require.Equal(t, []byte{0x06}, result)

	match, result = parse.Uint8("255")
	require.True(t, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = parse.Uint8("256")
	require.False(t, match)

	match, result = parse.Uint8("65535")
	require.False(t, match)

	match, result = parse.Uint8("x")
	require.False(t, match)

}

func TestUint16(t *testing.T) {
	match, result := parse.Uint16("6")
	require.True(t, match)
	require.Equal(t, []byte{0x06, 0x00}, result)

	match, result = parse.Uint16("255")
	require.True(t, match)
	require.Equal(t, []byte{0xff, 0x00}, result)

	match, result = parse.Uint16("65535")
	require.True(t, match)
	require.Equal(t, []byte{0xff, 0xff}, result)

	match, result = parse.Uint16("127236748")
	require.False(t, match)

	match, result = parse.Uint16("h")
	require.False(t, match)
}

func TestChar(t *testing.T) {
	match, result := parse.Char("'c'")
	require.True(t, match)
	require.Equal(t, []byte{0x63}, result)

	match, result = parse.Char("'C'")
	require.True(t, match)
	require.Equal(t, []byte{0x43}, result)

	match, result = parse.Char("' '")
	require.True(t, match)
	require.Equal(t, []byte{0x20}, result)

	match, result = parse.Char("'''")
	require.True(t, match)
	require.Equal(t, []byte{0x27}, result)

	match, result = parse.Char("'9'")
	require.True(t, match)
	require.Equal(t, []byte{0x39}, result)

	match, result = parse.Char("'")
	require.False(t, match)

	match, result = parse.Char("")
	require.False(t, match)

	match, result = parse.Char("'😀'")
	require.False(t, match)
}
