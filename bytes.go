package hackutils

import "log"

// XorBytes performs a binary XOR between two byte slices of the same length.
// If the length is different it fails. Returns a new byte slice with the
// result.
func XorBytes(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		log.Fatalf("len(b1) == %v, len(b2) == %v", len(b1), len(b2))
	}
	out := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		out[i] = b1[i] ^ b2[i]
	}
	return out
}

// XorBytesWithVal performs a binary XOR between every byte in bs with b.
// Returns a new byte slice with the result.
func XorBytesWithVal(bs []byte, b byte) []byte {
	out := make([]byte, len(bs))
	for i := 0; i < len(bs); i++ {
		out[i] = bs[i] ^ b
	}
	return out
}
