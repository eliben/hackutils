package hackutils

import (
	"testing"
)

// To compare the numbers generated, I dowloaded the original MT code from
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/CODES/mt19937ar.c,
// compiled it and ran with the same seed. The comparison here of several
// early and relatively late generated numbers make it very unlikely that my
// implementation is wrong. The state size is 624 and I'm generating 1000
// numbers which ensures that twist() was called.
func TestMT(t *testing.T) {
	// Use seed 100, which was given to the reference implementation too.
	ref := make(map[int]uint32)
	ref[0] = 2333906440
	ref[1] = 2882591512
	ref[2] = 1195587395
	ref[997] = 580208846
	ref[998] = 1302567035
	ref[999] = 2696579586

	mt := NewMT19937(100)
	numSaw := 0

	for i := 0; i < 1000; i++ {
		n := mt.Next()
		if refValue, ok := ref[i]; ok {
			if refValue != n {
				t.Errorf("for index %v: got %v, want %v", i, n, refValue)
			}
			numSaw++
		}
	}

	if numSaw != len(ref) {
		t.Errorf("saw %v values, want %v", numSaw, len(ref))
	}
}
