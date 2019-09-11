package hackutils

import (
	"log"
	"math"
)

// KLDivergence computes the Kullback-Leibler divergence between two probability
// distributions p and q. p and q are expected to have the same length and
// contain numbers that sum up to 1.0; The KL divergence is close to 0 if the
// two distributions are deemed "similar" and gets larger the more different
// they are.
// See https://en.wikipedia.org/wiki/Kullback%E2%80%93Leibler_divergence
func KLDivergence(p, q []float64) float64 {
	if len(p) != len(q) {
		log.Fatalf("len(p) == %v, len(q) == %v", len(p), len(q))
	}
	var total float64
	for i := 0; i < len(p); i++ {
		if p[i] > 0 {
			total += p[i] * math.Log(q[i]/p[i])
		}
	}
	return -total
}
