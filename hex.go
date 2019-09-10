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

// XorBytes performs a binary XOR between two byte slices of the same length.
// If the length is different it fails.
func XorBytes(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		log.Fatalf("len(b1) == %v != len(b2) == %v", len(b1), len(b2))
	}
	out := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		out[i] = b1[i] ^ b2[i]
	}
	return out
}
