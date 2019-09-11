package hackutils

import (
	"encoding/hex"
	"log"
)

// HexToBytes converts hexadecimal encoded bytes to raw bytes, failing the
// program in case of errors. For example, []byte{"aa18"} is converted to
// []byte{0xaa, 0x18}.
func HexToBytes(h []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(h)))
	n, err := hex.Decode(dst, h)
	if err != nil {
		log.Fatal(err)
	}
	if len(dst) != n {
		log.Fatalf("len dst == %v, d == %v", len(dst), n)
	}
	return dst
}
