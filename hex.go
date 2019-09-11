package hackutils

import (
	"encoding/hex"
	"log"
	"math"
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

// TODO: This and the Xors don't belong in hex.go -- split it up

// KLDivergence computes the Kullback-Leibler divergence between two probability
// distributions p and q. p and q are expected to have the same length and
// contain numbers that sum up to 1.0; The KL divergence is close to 0 if the
// two distributions are deemed "similar" and gets larger the more different
// they are.
// See https://en.wikipedia.org/wiki/Kullback%E2%80%93Leibler_divergence
func KLDivergence(p, q []float64) float64 {
	if len(p) != len(q) {
		log.Fatalf("len(p) == %v, len(q) == %v", len(p), len(q))
	}
	var total float64
	for i := 0; i < len(p); i++ {
		if p[i] > 0 {
			total += p[i] * math.Log(q[i]/p[i])
		}
	}
	return -total
}
