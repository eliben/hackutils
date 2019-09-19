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

func TestStripPKCS7Padding(t *testing.T) {
	var tests = []struct {
		data   []byte
		hasErr bool
		result []byte
	}{
		{[]byte{10, 20, 1}, false, []byte{10, 20}},
		{[]byte{10, 20, 0}, true, nil},
		{[]byte{10, 2, 2}, false, []byte{10}},
		{[]byte{1, 3, 3, 3}, false, []byte{1}},
		{[]byte{10, 20, 2}, true, nil},
		{[]byte{10, 20, 20}, true, nil},
		{[]byte{10, 20, 4, 4, 4, 4}, false, []byte{10, 20}},
		{[]byte{1, 2, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16}, false, []byte{1, 2}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.data)
		t.Run(testname, func(t *testing.T) {
			result, err := StripPKCS7Padding(tt.data)
			if err != nil {
				if !tt.hasErr {
					t.Errorf("err %v, expected to pass", err)
				} else {
					return
				}
			}
			if tt.hasErr {
				// err == nil, so this isn't expected.
				t.Errorf("expected err, got %v", tt.result)
			}
			if bytes.Compare(result, tt.result) != 0 {
				t.Errorf("got %v, want %v", result, tt.result)
			}
		})
	}
}

func TestPKCS7Roundtrip(t *testing.T) {
	// Padding is done to blocks of 16. Generate all blocks of size 1..31
	// and make sure their padding works round-trip.
	var tests [][]byte
	for i := 1; i <= 31; i++ {
		d := make([]byte, i)
		for c := 0; c < i; c++ {
			d[c] = byte(c) + 1
		}
		tests = append(tests, d)
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt)
		t.Run(testname, func(t *testing.T) {
			padding := PKCS7Padding(tt, 16)
			newdata := append(tt, padding...)
			nopad, err := StripPKCS7Padding(newdata)
			if err != nil {
				t.Error(err)
			}
			if bytes.Compare(nopad, tt) != 0 {
				t.Errorf("got %v, want %v", nopad, tt)
			}
		})
	}
}
