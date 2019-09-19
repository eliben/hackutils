package hackutils

import (
	"bytes"
	"testing"
)

func TestBase64ToBytes(t *testing.T) {
	var tests = []struct {
		in  []byte
		out []byte
	}{
		{[]byte(""), []byte{}},
		{[]byte("aGVsbG8="), []byte("hello")},
		{[]byte("aXMgaXQgbWUgeW91J3JlIGxvb2tpbmcgZm9yPw=="), []byte("is it me you're looking for?")},
		{[]byte("AQIDBA=="), []byte{0x01, 0x02, 0x03, 0x04}},
	}
	for _, tt := range tests {
		t.Run(string(tt.in), func(t *testing.T) {
			result := Base64ToBytes(tt.in)
			if bytes.Compare(result, tt.out) != 0 {
				t.Errorf("got %v, want %v", result, tt.out)
			}
		})
	}
}
