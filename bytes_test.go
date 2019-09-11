package hackutils

import (
	"bytes"
	"testing"
)

func TestXorBytes(t *testing.T) {
	b1 := []byte{0x21, 0x7a, 0xff}
	b2 := []byte{0xff, 0xf0, 0xab}
	expected := []byte{0xde, 0x8a, 0x54}
	result := XorBytes(b1, b2)
	if bytes.Compare(result, expected) != 0 {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestXorBytesWithVal(t *testing.T) {
	bs := []byte{0x78, 0x33, 0x51, 0xac}
	var b byte = 0x55
	expected := []byte{0x2d, 0x66, 0x04, 0xf9}
	result := XorBytesWithVal(bs, b)
	if bytes.Compare(result, expected) != 0 {
		t.Errorf("got %v, want %v", result, expected)
	}
}
