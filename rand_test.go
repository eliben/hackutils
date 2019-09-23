package hackutils

import (
	"fmt"
	"testing"
)

func TestMTSeed(t *testing.T) {
	mt := NewMT19937(100)
	for i := 0; i < 5; i++ {
		fmt.Println(mt.Next())
	}
}
