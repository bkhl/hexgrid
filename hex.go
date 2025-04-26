package hexgrid

import (
	"fmt"
	"math"
	"slices"
)

type Direction int

const (
	DirectionSE = iota
	DirectionNE
	DirectionN
	DirectionNW
	DirectionSW
	DirectionS
)

var directions = []Hex{
	NewHex(1, 0),
	NewHex(1, -1),
	NewHex(0, -1),
	NewHex(-1, 0),
	NewHex(-1, +1),
	NewHex(0, +1),
}

// hex describes a regular hexagon with Cube Coordinates (although the S coordinate is computed on the constructor)
// It's also easy to reference them as axial (trapezoidal coordinates):
// - R represents the vertical axis
// - Q the diagonal one
// - S can be ignored
// For additional reference on these coordinate systems: http://www.redblobgames.com/grids/hexagons/#coordinates
//
//	         _ _
//	       /     \
//	  _ _ /(0,-1) \ _ _
//	/     \  -R   /     \
//
// /(-1,0) \ _ _ /(1,-1) \
// \  -Q   /     \       /
//
//	\ _ _ / (0,0) \ _ _ /
//	/     \       /     \
//
// /(-1,1) \ _ _ / (1,0) \
// \       /     \  +Q   /
//
//	\ _ _ / (0,1) \ _ _ /
//	      \  +R   /
//	       \ _ _ /
type Hex struct {
	Q int // x axis
	R int // y axis
	S int // z axis
}

func NewHex(q, r int) Hex {
	return Hex{Q: q, R: r, S: -q - r}
}

func (h Hex) String() string {
	return fmt.Sprintf("(%d,%d)", h.Q, h.R)
}

// Adds another hexagon
func (h Hex) Add(o Hex) Hex {
	return NewHex(h.Q+o.Q, h.R+o.R)
}

// Subtracts another hexagon
func (h Hex) Subtract(o Hex) Hex {
	return NewHex(h.Q-o.Q, h.R-o.R)
}

// Scales an hexagon by a k factor. If factor k is 1 there's no change
func (h Hex) Scale(k int) Hex {
	return NewHex(h.Q*k, h.R*k)
}

func (h Hex) Length() int {
	return int((math.Abs(float64(h.Q)) + math.Abs(float64(h.R)) + math.Abs(float64(h.S))) / 2.)
}

func (h Hex) Distance(o Hex) int {
	sub := h.Subtract(o)
	return sub.Length()
}

// Returns the neighbor hexagon at a certain direction
// TODO: Not sure if I like this approach
func (h Hex) Neighbor(direction Direction) Hex {
	return h.Add(directions[direction])
}

// Returns all neighboring hexagons
func (h Hex) Neighbors() []Hex {
	neighbors := make([]Hex, len(directions))
	for i, d := range directions {
		neighbors[i] = h.Add(d)
	}
	return neighbors
}

// Returns the slice of hexagons that exist on a line that goes from the hexagon to another
// TODO: Name doesn't seem to fit
func (h Hex) LineDraw(o Hex) []Hex {
	hexLerp := func(a FractionalHex, b FractionalHex, t float64) FractionalHex {
		return NewFractionalHex(a.Q*(1-t)+b.Q*t, a.R*(1-t)+b.R*t)
	}

	N := h.Distance(o)

	// Sometimes the hexLerp will output a point that’s on an edge.
	// On some systems, the rounding code will push that to one side or the other,
	// somewhat unpredictably and inconsistently.
	// To make it always push these points in the same direction, add an “epsilon” value to a.
	// This will “nudge” things in the same direction when it’s on an edge, and leave other points unaffected.

	a_nudge := NewFractionalHex(float64(h.Q)+0.000001, float64(h.R)+0.000001)
	b_nudge := NewFractionalHex(float64(o.Q)+0.000001, float64(o.R)+0.000001)

	results := make([]Hex, 0)
	step := 1. / math.Max(float64(N), 1)

	for i := 0; i <= N; i++ {
		results = append(results, hexLerp(a_nudge, b_nudge, step*float64(i)).Round())
	}
	return results
}

// Returns the set of hexagons around the hexagon for a given radius
func (h Hex) Range(r int) []Hex {
	var results = make([]Hex, 0)
	if r >= 0 {
		for dx := -r; dx <= r; dx++ {
			for dy := math.Max(float64(-r), float64(-dx-r)); dy <= math.Min(float64(r), float64(-dx+r)); dy++ {
				results = append(results, h.Add(NewHex(int(dx), int(dy))))
			}
		}
	}
	return results
}

// Determines if a given hexagon is visible from the hexagon, taking into consideration a set of blocking hexagons
func (h Hex) HasLineOfSight(target Hex, blocking []Hex) bool {
	contains := func(s []Hex, e Hex) bool {
		return slices.Contains(s, e)
	}

	for _, hexOnLine := range h.LineDraw(target) {
		if contains(blocking, hexOnLine) {
			return false
		}
	}

	return true
}

// Returns the list of hexagons that are visible from the hexagon
func (h Hex) FieldOfView(candidates []Hex, blocking []Hex) []Hex {
	results := make([]Hex, 0)
	for _, candidate := range candidates {
		distance := h.Distance(candidate)
		if len(blocking) == 0 || distance <= 1 || h.HasLineOfSight(candidate, blocking) {
			results = append(results, candidate)
		}
	}
	return results
}

// Returns the set of hexagons that form a rectangle with the specified width and height
func RectangleGrid(width, height int) []Hex {
	results := make([]Hex, 0)
	for q := range width {
		qOffset := int(math.Floor(float64(q) / 2.))
		for r := -qOffset; r < height-qOffset; r++ {
			results = append(results, NewHex(q, r))
		}
	}
	return results
}
