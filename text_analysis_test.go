package hackutils

import (
	"testing"
)

func TestFrequencyDistributionScore(t *testing.T) {
	// Texts ordered by how close to English they are, from closest to farthest.
	ts := [][]byte{
		[]byte("this a piece of standard english text"),
		[]byte("this a slightly less standard zkzkzkk"),
		[]byte("cxvpqewkw ssylkvjcxyq sadfpqlkzxhgfri"),
		[]byte("xc@#$dcf,.xs@24309u8@#08753 .][/12aaa"),
	}

	var scores []float64
	for _, t := range ts {
		score := FrequencyDistributionScore(t)
		scores = append(scores, score)
	}

	for i := 1; i < len(scores); i++ {
		if scores[i] < scores[i-1] {
			t.Errorf("[%d] %f < [%d] %f", i, scores[i], i-1, scores[i-1])
		}
	}
}
