package hackutils

import (
	"bytes"
	"testing"
)

func TestHexToBytes(t *testing.T) {
	var tests = []struct {
		in  []byte
		out []byte
	}{
		{[]byte(""), []byte{}},
		{[]byte("ab"), []byte{0xab}},
		{[]byte("ab01"), []byte{0xab, 0x01}},
		{[]byte("07214fa893"), []byte{0x07, 0x21, 0x4f, 0xa8, 0x93}},
	}
	for _, tt := range tests {
		t.Run(string(tt.in), func(t *testing.T) {
			result := HexToBytes(tt.in)
			if bytes.Compare(result, tt.out) != 0 {
				t.Errorf("got %v, want %v", result, tt.out)
			}
		})
	}
}

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

func TestKLDivergence(t *testing.T) {
	as := []float64{0.33, 0.33, 0.1, 0.23}

	// Intuitively, these are ordered from the "closest to as" to "farthest from
	// as".
	bs := []float64{0.32, 0.34, 0.09, 0.24}
	cs := []float64{0.27, 0.39, 0.08, 0.25}
	ds := []float64{0.14, 0.51, 0.25, 0.1}
	es := []float64{0.1, 0.2, 0.3, 0.4}

	divBs := KLDivergence(as, bs)
	divCs := KLDivergence(as, cs)
	divDs := KLDivergence(as, ds)
	divEs := KLDivergence(as, es)

	if divBs > divCs {
		t.Errorf("got divBs > divCs")
	}
	if divCs > divDs {
		t.Errorf("got divCs > divDs")
	}
	if divDs > divEs {
		t.Errorf("got divDs > divEs")
	}
}
