package hexgrid

import "math"

// fractionHex provides a more precise representation for hexagons when precision is required.
// It's also represented in Cube Coordinates
type FractionalHex struct {
	q float64
	r float64
	s float64
}

func NewFractionalHex(q, r float64) FractionalHex {
	h := FractionalHex{q: q, r: r, s: -q - r}
	return h
}

// Rounds a FractionalHex to a Regular Hex
func (h FractionalHex) Round() Hex {
	roundToInt := func(a float64) int {
		if a < 0 {
			return int(a - 0.5)
		}
		return int(a + 0.5)
	}

	q := roundToInt(h.q)
	r := roundToInt(h.r)
	s := roundToInt(h.s)

	q_diff := math.Abs(float64(q) - h.q)
	r_diff := math.Abs(float64(r) - h.r)
	s_diff := math.Abs(float64(s) - h.s)

	if q_diff > r_diff && q_diff > s_diff {
		q = -r - s
	} else if r_diff > s_diff {
		r = -q - s
	} else {
		s = -q - r
	}
	return Hex{q, r, s}
}
