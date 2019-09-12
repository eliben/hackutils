package hackutils

import (
	"bufio"
	"encoding/hex"
	"io"
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

// HexReaderToBytes converts all hexadecimal-encoded data in r to raw bytes.
// The input has to consist of (possibly whitespace-separated) hexadecimal
// daat only - for example a file with some hexa data per line (newline
// separated). The returned byte buffer is the concatenation of all input
// found in the reader.
func HexReaderToBytes(r io.Reader) []byte {
	var out []byte
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		out = append(out, HexToBytes(scanner.Bytes())...)
	}
	return out
}
