package hackutils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
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

func TestXorWithRepeatingMask(t *testing.T) {
	mask := []byte{0xff, 0xf0, 0x0f}
	bs := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77}

	// full bs: has one leftover byte
	result1 := XorWithRepeatingMask(bs, mask)
	expected1 := []byte{0xee, 0xd2, 0x3c, 0xbb, 0xa5, 0x69, 0x88}
	if bytes.Compare(result1, expected1) != 0 {
		t.Errorf("got %v, want %v", result1, expected1)
	}

	// just 6 bytes of bs: no leftover
	result2 := XorWithRepeatingMask(bs[:6], mask)
	expected2 := expected1[:6]
	if bytes.Compare(result2, expected2) != 0 {
		t.Errorf("got %v, want %v", result2, expected2)
	}

	// bs shorter than mask
	result3 := XorWithRepeatingMask(bs[:2], mask)
	expected3 := expected1[:2]
	if bytes.Compare(result3, expected3) != 0 {
		t.Errorf("got %v, want %v", result3, expected3)
	}
}

func TestHammingDistance(t *testing.T) {
	var tests = []struct {
		b1     []byte
		b2     []byte
		expect int64
	}{
		{[]byte{0x01, 0x02, 0x03}, []byte{0x01, 0x02, 0x03}, 0},
		{[]byte{0x01, 0x02, 0x03}, []byte{0x03, 0x02, 0x03}, 1},
		{[]byte{0x01, 0x02, 0x03}, []byte{0x02, 0x03, 0x04}, 6},
		{[]byte("this is a test"), []byte("wokka wokka!!!"), 37},
		{bytes.Repeat([]byte{0xff, 0x55}, 50), bytes.Repeat([]byte{0xff, 0xaa}, 50), 400},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := HammingDistance(tt.b1, tt.b2)
			if result != tt.expect {
				t.Errorf("got %v, want %v", result, tt.expect)
			}
		})
	}
}

func TestRandBytes(t *testing.T) {
	b := RandBytes(32)
	if len(b) != 32 {
		t.Errorf("len(b) == %v, want 32", len(b))
	}

	// Sanity check that the bytes returned are not all equal (this *can* happen
	// randomly but the chances are negligibly low).
	foundDiff := false
	for i := 1; i < len(b); i++ {
		if b[i] != b[i-1] {
			foundDiff = true
			break
		}
	}
	if !foundDiff {
		t.Errorf("all bytes in b are equal, something's wrong")
	}
}

func TestDumpBytesGrid(t *testing.T) {
	bs := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := `
       0  1  2  3 
     ------------
   0| 01 02 03 04 
   4| 05 06 07 08 
   8| 09 0A `

	result := DumpBytesGrid(bs, 4)
	if strings.TrimSpace(expected) != strings.TrimSpace(result) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestPackUint64LE(t *testing.T) {
	var tests = []struct {
		in  uint64
		out []byte
	}{
		{0, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, []byte{1, 0, 0, 0, 0, 0, 0, 0}},
		{258, []byte{2, 1, 0, 0, 0, 0, 0, 0}},
		{0x202122, []byte{34, 33, 32, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			result := PackUint64LE(tt.in)
			if bytes.Compare(result, tt.out) != 0 {
				t.Errorf("got %v, want %v", result, tt.out)
			}
		})
	}
}
