package hexgrid

import "math"

type Point struct {
	X float64
	Y float64
}

type Layout struct {
	Orientation orientation
	Size        Point // multiplication factor relative to the canonical hexagon, where the points are on a unit circle
	Origin      Point // center Point for hexagon 0,0
}

type orientation struct {
	f0, f1, f2, f3, b0, b1, b2, b3, startAngle float64
}

var OrientationPointy orientation = orientation{math.Sqrt(3.), math.Sqrt(3.) / 2., 0., 3. / 2., math.Sqrt(3.) / 3., -1. / 3., 0., 2. / 3., 0.5}

var OrientationFlat orientation = orientation{3. / 2., 0., math.Sqrt(3.) / 2., math.Sqrt(3.), 2. / 3., 0., -1. / 3., math.Sqrt(3.) / 3., 0.}

// HexToPixel returns the center pixel for a given hexagon an a certain Layout
func HexToPixel(l Layout, h Hex) Point {

	M := l.Orientation
	size := l.Size
	origin := l.Origin
	x := (M.f0*float64(h.Q) + M.f1*float64(h.R)) * size.X
	y := (M.f2*float64(h.Q) + M.f3*float64(h.R)) * size.Y
	return Point{x + origin.X, y + origin.Y}
}

// PixelToHex returns the corresponding hexagon axial coordinates for a given pixel on a certain Layout
func PixelToHex(l Layout, p Point) FractionalHex {

	M := l.Orientation
	size := l.Size
	origin := l.Origin

	pt := Point{(p.X - origin.X) / size.X, (p.Y - origin.Y) / size.Y}
	q := M.b0*pt.X + M.b1*pt.Y
	r := M.b2*pt.X + M.b3*pt.Y
	return FractionalHex{q, r, -q - r}
}

func HexCornerOffset(l Layout, c int) Point {

	M := l.Orientation
	size := l.Size
	angle := 2. * math.Pi * (M.startAngle - float64(c)) / 6.
	return Point{size.X * math.Cos(angle), size.Y * math.Sin(angle)}
}

// Gets the corners of the hexagon for the given Layout, starting at the E vertex and proceeding in a CCW order
func HexagonCorners(l Layout, h Hex) []Point {

	corners := make([]Point, 0)
	center := HexToPixel(l, h)

	for i := range 6 {
		offset := HexCornerOffset(l, i)
		corners = append(corners, Point{center.X + offset.X, center.Y + offset.Y})
	}
	return corners
}
