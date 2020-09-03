package parse_test

import (
	"testing"

	"github.com/snarechops/assembler/parse"
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

	match, result = parse.LDA("LDA 0xff12")
	require.Equal(t, `LDA`, match)
	require.Equal(t, []byte{0xa9, 0x12, 0xff}, result)
}

func TestLDX(t *testing.T) {
	match, result := parse.LDX("LDX $0xff12")
	require.Equal(t, `LDX`, match)
	require.Equal(t, []byte{0xae, 0x12, 0xff}, result)

	match, result = parse.LDX("LDX $0xff12,y")
	require.Equal(t, `LDX`, match)
	require.Equal(t, []byte{0xbe, 0x12, 0xff}, result)

	match, result = parse.LDX("LDX 0xff12")
	require.Equal(t, `LDX`, match)
	require.Equal(t, []byte{0xa2, 0x12, 0xff}, result)

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

	match, result = parse.LDY("LDY 0xff12")
	require.Equal(t, `LDY`, match)
	require.Equal(t, []byte{0xa0, 0x12, 0xff}, result)

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
