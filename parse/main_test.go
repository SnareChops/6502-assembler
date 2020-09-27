package parse_test

import (
	"testing"

	"github.com/snarechops/6502-assembler/parse"
	"github.com/stretchr/testify/require"
)

func TestLine(t *testing.T) {
	result := parse.Line("label:", 0)
	require.Len(t, result, 0)

	result = parse.Line("  other:    ", 0)
	require.Len(t, result, 0)

	result = parse.Line("LDA $0xff12", 0)
	require.Len(t, result, 3)

	result = parse.Line("  LDA $0xff12    ", 0)
	require.Len(t, result, 3)

	result = parse.Line("", 0)
	require.Len(t, result, 0)

	result = parse.Line("// comment", 0)
	require.Len(t, result, 0)

	require.Panics(t, func() { parse.Line("sfoksdf wefoh sdf", 0) }, 0)
}
