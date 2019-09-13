package hackutils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPKCS7Pad(t *testing.T) {
	var tests = []struct {
		dataLen  int
		blockLen int
		padding  []byte
	}{
		{7, 4, []byte{1}},
		{10, 4, []byte{2, 2}},
		{17, 4, []byte{3, 3, 3}},
		{100, 4, []byte{4, 4, 4, 4}},
		{101, 16, []byte{11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("pad-%d-to-%d", tt.dataLen, tt.blockLen)
		t.Run(testname, func(t *testing.T) {
			result := PKCS7Padding(bytes.Repeat([]byte{0}, tt.dataLen), tt.blockLen)
			if bytes.Compare(result, tt.padding) != 0 {
				t.Errorf("got %v, want %v", result, tt.padding)
			}
		})
	}
}
