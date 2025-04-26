# Hexgrid

This is a fork of <https://github.com/Laminator42/hexgrid> which in turn was
forked from <https://github.com/pmcxs/hexgrid>. It's based on the algorithms
described at <http://www.redblobgames.com/grids/hexagons/implementation.html>.

## Installation

    go get github.com/bkhl/hexgrid

## Usage

#### Importing

```go
import "github.com/bkhl/hexgrid"
```

### Examples

#### Creating hexagons

```go
hexagonA := NewHex(1,2) //at axial coordinates Q=1 R=2
hexagonB := NewHex(2,3) //at axial coordinates Q=2 R=3
```

#### Measuring the distance (in hexagons) between two hexagons

```go
distance := hexagonA.Distance(hexagonB)
```

#### Getting the array of hexagons on the path between two hexagons

```go
origin := NewHex(10,20)
destination := NewHex(30,40)
path := origin.LineDraw(destination)
```

#### Creating a layout

```go
origin := point {0,0}     // The coordinate that corresponds to the center of hexagon 0,0
size := point {100, 100}  // The length of an hexagon side => 100
layout: = layout{size, origin, orientationFlat}
```

#### Obtaining the pixel that corresponds to a given hexagon

```go
hex := NewHex(1,0)
pixel := HexToPixel(layout,hex)  // Pixel that corresponds to the center of hex 1,0 (in the given layout)
```

#### Obtaining the hexagon that contains the given pixel (and rounding it)

```go
point := point {10,20}
hex := PixelToHex(layout, point).Round()
```

## Credits

* Jannik Bach (<https://github.com/Laminator42>)
* Pedro Sousa (<https://github.com/pmcxs>)
* Red Blob Games (<http://www.redblobgames.com/grids/hexagons/implementation.html>)

## License

MIT
