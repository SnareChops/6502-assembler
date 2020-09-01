package conv

import "encoding/binary"

// Uint16 Convert uint16 to [2]byte in little endian format
func Uint16(num uint16) []byte {
	result := make([]byte, 2)
	binary.LittleEndian.PutUint16(result, num)
	return result
}
