package hackutils

import (
	"bytes"
	"fmt"
	"log"
)

// PKCS7Padding returns the padding required to extend data to a multiple of
// blockLen, following the PKCS#7 scheme from RFC 5652.
func PKCS7Padding(data []byte, blockLen int) []byte {
	if blockLen < 2 {
		log.Fatalf("expect blockLen < 0, got %v", blockLen)
	}
	padLen := blockLen - len(data)%blockLen
	return bytes.Repeat([]byte{byte(padLen)}, padLen)
}

// StripPKCS7Padding strips PKCS#7 padding from data. Returns a slice of the
// data without padding (likely to alias data). Error is returned when the
// padding is incorrect.
func StripPKCS7Padding(data []byte) ([]byte, error) {
	// Look at the last byte of padding. For N, expect the last N bytes to be
	// equal to N and strip them. 0 is invalid padding.
	npad := data[len(data)-1]
	if npad == 0 || len(data) <= int(npad) {
		return nil, fmt.Errorf("padding error: len(data)=%v, npad=%d", len(data), npad)
	}

	// Check that all padding bytes are correct.
	for i := 0; i < int(npad); i++ {
		if data[len(data)-1-i] != npad {
			return nil, fmt.Errorf("padding error: data[%d]=%d, npad=%d", len(data)-1-i, data[len(data)-1-i], npad)
		}
	}

	return data[:len(data)-int(npad)], nil
}
