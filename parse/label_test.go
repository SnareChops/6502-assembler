package parse_test

import (
	"fmt"
	"testing"

	"github.com/snarechops/assembler/parse"
	"github.com/stretchr/testify/require"
)

func TestLabel(t *testing.T) {
	match, err := parse.Label("test:", 0)
	require.True(t, match)
	require.Nil(t, err)
	require.Equal(t, []byte{0x00, 0x00}, parse.GetLabel("test"))

	match, err = parse.Label("t_16_ahs:", 12)
	require.True(t, match)
	require.Nil(t, err)
	require.Equal(t, []byte{0x0c, 0x00}, parse.GetLabel("t_16_AHs"))

	match, err = parse.Label("test:", 42)
	require.False(t, match)
	require.NotNil(t, err)
	require.Equal(t, "Duplicate label 'test'", fmt.Sprint(err))

	match, err = parse.Label("test", 64)
	require.False(t, match)
	require.Nil(t, err)

	match, err = parse.Label("222:", 74)
	require.False(t, match)
	require.NotNil(t, err)
	require.Equal(t, "Label must start with a letter '222'", fmt.Sprint(err))

	parse.ClearLabels()
}
