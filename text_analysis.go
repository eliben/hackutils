package hackutils

import (
	"fmt"
	"math"
)

// Probability distribution for the 26 alphabetic letters in English.
var masterFreq = []float64{
	0.0817, 0.0149, 0.0278, 0.0425, 0.1270, 0.0223, // a-f
	0.0201, 0.0609, 0.0698, 0.0015, 0.0077, 0.0402, // g-l
	0.0241, 0.0675, 0.0752, 0.0193, 0.0009, 0.0599, // m-r
	0.0633, 0.0906, 0.0277, 0.0098, 0.0236, 0.0015, // s-x
	0.0197, 0.0007, // y-z
}

// FrequencyDistributionScore computes a heuristic score for the frequency
// distribution of letters in bs, relative to masterFreq. It also applies some
// heurstics to determine if the text looks like English at all. The lower the
// score, the closer to English it looks.
func FrequencyDistributionScore(bs []byte) float64 {
	hist := make([]int64, 26)

	var total int
	var nonprintable int

	for _, b := range bs {
		if b >= 65 && b <= 90 {
			b += 32
		}
		// Now b is either in the lowercase range, or uncounted.
		if b >= 97 && b <= 122 {
			hist[b-97]++
			total++
		}

		if b < 9 || (b > 13 && b < 32) {
			fmt.Println(b)
			nonprintable++
		}
	}
	fmt.Println("**", len(bs), total, nonprintable)

	// Here total is the number of alpha chars in the string, and len(bs) is the
	// number of all chars. Heuristically we don't expect more than 20-25% of
	// the chars to be non-alpha. We'll take a safety margin to 33% and will be
	// adding a penalty to the score for each char beyond that.
	penalty := 0.0
	badchars := len(bs) - total
	if badchars*3 >= len(bs) {
		penalty = 0.15 * (float64(badchars) - math.Floor(float64(len(bs))/3.0))
	}

	// We also add a penalty of 0.15 for each non-printable character
	penalty += 0.15 * float64(nonprintable)

	// Here hist is a histogram of the appearance of the 26 letters in bs. total
	// is the sum of the histogram. We normalize it to a probability distribution.
	prob := make([]float64, 26)
	for i, h := range hist {
		prob[i] = float64(h) / float64(total)
	}

	// Finally, we compute the KL divergence of prob compared to masterFreq.
	return penalty + KLDivergence(prob, masterFreq)
}
