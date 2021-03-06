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

func TestHexReaderToBytes(t *testing.T) {
	buf := bytes.NewBuffer([]byte("ab0702\n0409\nab"))
	expected := []byte{0xab, 0x07, 0x02, 0x04, 0x09, 0xab}
	result := HexReaderToBytes(buf)
	if bytes.Compare(result, expected) != 0 {
		t.Errorf("got %v, want %v", result, expected)
	}
}
