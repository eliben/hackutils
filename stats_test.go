package hackutils

import "testing"

func TestKLDivergence(t *testing.T) {
	as := []float64{0.33, 0.33, 0.1, 0.23}

	// Intuitively, these are ordered from the "closest to as" to "farthest from
	// as".
	bs := []float64{0.32, 0.34, 0.09, 0.24}
	cs := []float64{0.27, 0.39, 0.08, 0.25}
	ds := []float64{0.14, 0.51, 0.25, 0.1}
	es := []float64{0.1, 0.2, 0.3, 0.4}

	divBs := KLDivergence(as, bs)
	divCs := KLDivergence(as, cs)
	divDs := KLDivergence(as, ds)
	divEs := KLDivergence(as, es)

	if divBs > divCs {
		t.Errorf("got divBs > divCs")
	}
	if divCs > divDs {
		t.Errorf("got divCs > divDs")
	}
	if divDs > divEs {
		t.Errorf("got divDs > divEs")
	}
}
