package hackutils

import (
	"bytes"
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
