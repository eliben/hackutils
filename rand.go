package hackutils

// Implementation of the Mersenne Twister PRNG algorithm, following the
// constants and pseudocode from https://en.wikipedia.org/wiki/Mersenne_Twister,
// for the 32-bit version.
const mtW = 32
const mtN = 624
const mtM = 397
const mtR = 31
const mtF = 1812433253
const mtA = 0x9908B0DF
const mtU = 11
const mtD = 0xFFFFFFFF
const mtS = 7
const mtB = 0x9D2C5680
const mtT = 15
const mtC = 0xEFC60000
const mtL = 18

const mtLowerMask uint32 = (1 << mtR) - 1
const mtUpperMask uint32 = ^mtLowerMask

type MT19937 struct {
	state [mtN]uint32
	index uint32
}

// NewMT19937 creates a new MT PRNG with the given seed.
func NewMT19937(seed uint32) *MT19937 {
	mt := new(MT19937)
	mt.state[0] = seed
	mt.index = mtN

	for i := 1; i < mtN; i++ {
		xn1 := mt.state[i-1]
		mt.state[i] = mtF*(xn1^(xn1>>(mtW-2))) + uint32(i)
	}
	return mt
}

func (mt *MT19937) Next() uint32 {
	if mt.index >= mtN {
		mt.twist()
	}

	y := mt.state[mt.index]
	y ^= ((y >> mtU) & mtD)
	y ^= ((y << mtS) & mtB)
	y ^= ((y << mtT) & mtC)
	y ^= (y >> mtL)

	mt.index++
	return y
}

func (mt *MT19937) twist() {
	for i := 0; i < mtN; i++ {
		x := (mt.state[i] & mtUpperMask) + (mt.state[(i+1)%mtN] & mtLowerMask)
		xA := x >> 1
		if x%2 != 0 {
			xA ^= mtA
		}
		mt.state[i] = mt.state[(i+mtM)%mtN] ^ xA
	}
	mt.index = 0
}
