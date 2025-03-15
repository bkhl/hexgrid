package hexgrid

import "math"

// fractionHex provides a more precise representation for hexagons when precision is required.
// It's also represented in Cube Coordinates
type FractionalHex struct {
	Q float64
	R float64
	S float64
}

func NewFractionalHex(q, r float64) FractionalHex {
	h := FractionalHex{Q: q, R: r, S: -q - r}
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

	q := roundToInt(h.Q)
	r := roundToInt(h.R)
	s := roundToInt(h.S)

	q_diff := math.Abs(float64(q) - h.Q)
	r_diff := math.Abs(float64(r) - h.R)
	s_diff := math.Abs(float64(s) - h.S)

	if q_diff > r_diff && q_diff > s_diff {
		q = -r - s
	} else if r_diff > s_diff {
		r = -q - s
	} else {
		s = -q - r
	}
	return Hex{q, r, s}
}
