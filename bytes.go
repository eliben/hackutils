package hackutils

import (
	"bytes"
	crand "crypto/rand"
	"log"
	"math/bits"
)

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

// XorWithRepeatingMask performs a binary XOR between bs and mask, with mask
// repeated enough times to cover the length of bs. For example, with a 10-byte
// bs and the mask "XYZ", bs will be XOR-ed with "XYZXYZXYZX". Note that just
// enough of the mask is taken to cover the leftovers.
func XorWithRepeatingMask(bs []byte, mask []byte) []byte {
	lb := len(bs)
	lm := len(mask)

	fullreps := lb / lm
	fullmask := bytes.Repeat(mask, fullreps)
	fullmask = append(fullmask, mask[:lb-len(fullmask)]...)

	return XorBytes(bs, fullmask)
}

// HammingDistance computes the hamming distance (number of bits that are
// different) between b1 and b2.
func HammingDistance(b1, b2 []byte) int64 {
	delta := XorBytes(b1, b2)
	total := 0
	for _, d := range delta {
		total += bits.OnesCount8(d)
	}
	return int64(total)
}

// RandBytes returns a new buffer with n random bytes.
func RandBytes(n int) []byte {
	b := make([]byte, n)
	_, err := crand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
