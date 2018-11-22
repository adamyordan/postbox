package postbox

import "encoding/binary"

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
