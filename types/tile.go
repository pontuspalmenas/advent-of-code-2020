package types

import (
	. "aoc"
	"fmt"
)

// A Tile_OLD wraps a Matrix, storing runes
type Tile_OLD struct {
	ID   int
	matrix *Matrix
}

func NewTile(id int, width, height int) Tile_OLD {
	m := NewMatrix(width, height)
	return Tile_OLD{ID: id, matrix: &m}
}

// Loads a Tile_OLD from string, e.g:
/*
Tile_OLD 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#
*/
func NewTileFromString_OLD(s string) Tile_OLD {
	lines := Lines(s)
	h := len(lines)-1 // Drop header
	w := len(lines[1]) // Check first row length
	tile := NewTile(Int(Regex(`Tile_OLD (\d+)`, lines[0])[0]), w, h)
	lines = lines[1:]
	for i := 0; i <= h-1; i++ {
		for j := 0; j <= w-1; j++ {
			tile.Set(Point{X: j, Y: i}, rune(lines[i][j]))
		}
	}
	return tile
}

func (t *Tile_OLD) Copy() *Tile_OLD {
	return &Tile_OLD{t.ID, t.matrix.Copy()}
}

func (t *Tile_OLD) Flip() {
	t.matrix.Flip()
}

// Rotates the tile 90 degrees clockwise
func (t *Tile_OLD) Rotate() {
	t.matrix.Rotate()
}

// Returns the state at point p
func (t *Tile_OLD) At(p Point) rune {
	return rune(t.matrix.At(p))
}

// Sets the state at point p
func (t *Tile_OLD) Set(p Point, v rune) {
	t.matrix.Set(p, int(v))
}

func (t *Tile_OLD) String() string {
	out := fmt.Sprintf("Tile_OLD %d:\n", t.ID)
	out += t.matrix.String()
	return out
}

func (t *Tile_OLD) checkBounds(p Point) {
	if p.X < 0 || p.X > t.Width() - 1 || p.Y < 0 || p.Y > t.Height() - 1 {
		panic(fmt.Sprintf("point out of bounds: %v", p))
	}
}

func (t *Tile_OLD) Width() int {
	return t.matrix.Width()
}

func (t *Tile_OLD) Height() int {
	return t.matrix.Height()
}

func (t *Tile_OLD) BorderTop() string {
	out := ""
	for i:=0; i<=t.Width()-1; i++ {
		out += string(t.At(Point{X: i}))
	}
	return out
}

func (t *Tile_OLD) BorderBottom() string {
	out := ""
	for i:=0; i<=t.Width()-1; i++ {
		out += string(t.At(Point{X: i, Y: t.Height()-1}))
	}
	return out
}

func (t *Tile_OLD) BorderLeft() string {
	out := ""
	for i:=0; i<=t.Width()-1; i++ {
		out += string(t.At(Point{Y: i}))
	}
	return out
}

func (t *Tile_OLD) BorderRight() string {
	out := ""
	for i:=0; i<=t.Width()-1; i++ {
		out += string(t.At(Point{X: t.Height()-1, Y: i}))
	}
	return out
}

func (t *Tile_OLD) Inner() string {
	s := ""
	for y:=1; y<t.Height(); y++ {
		for x:=1; x<=t.Width()-2;x++ {
			s += string(t.At(Point{X: x, Y: y}))
		}
	}
	return s
}

func (t *Tile_OLD) Column(row int) string {
	out := ""
	for i:=0; i <= t.Width()-1; i++ {
		out += string(t.At(Point{X: i, Y: row}))
	}
	return out
}

func (t *Tile_OLD) Variations() []Tile_OLD {
	var variations []Tile_OLD

	// Rotate three times, flip and rotate three times again
	for i:=0; i<2; i++ {
		variations = append(variations, *t.Copy())
		for j := 0; j < 3; j++ {
			t.Rotate()
			variations = append(variations, *t.Copy())
		}
		t.Flip()
	}
	return variations
}
