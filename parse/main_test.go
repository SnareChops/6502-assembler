package parse_test

import (
	"testing"

	"github.com/snarechops/assembler/parse"
	"github.com/stretchr/testify/require"
)

func TestLine(t *testing.T) {
	result := parse.Line("label:")
	require.Len(t, result, 0)

	result = parse.Line("  other:    ")
	require.Len(t, result, 0)

	result = parse.Line("LDA $0xff12")
	require.Len(t, result, 3)

	result = parse.Line("  LDA $0xff12    ")
	require.Len(t, result, 3)

	result = parse.Line("")
	require.Len(t, result, 0)

	result = parse.Line("// comment")
	require.Len(t, result, 0)

	require.Panics(t, func() { parse.Line("sfoksdf wefoh sdf") }, 0)
}
