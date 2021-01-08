package types

import (
	. "aoc"
	"fmt"
)

// A Tile wraps a Matrix, storing runes
type Tile struct {
	ID   int
	matrix *Matrix
}

func NewTile(id int, width, height int) Tile {
	m := NewMatrix(width, height)
	return Tile{ID: id, matrix: &m}
}

// Loads a Tile from string, e.g:
/*
Tile 2971:
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
func NewTileFromString(s string) Tile {
	lines := Lines(s)
	h := len(lines)-1 // Drop header
	w := len(lines[1]) // Check first row length
	tile := NewTile(Int(Regex(`Tile (\d+)`, lines[0])[0]), w, h)
	lines = lines[1:]
	for i := 0; i <= h-1; i++ {
		for j := 0; j <= w-1; j++ {
			tile.Set(Point{X: j, Y: i}, rune(lines[i][j]))
		}
	}
	return tile
}

func (t *Tile) Copy() *Tile {
	return &Tile{t.ID, t.matrix.Copy()}
}

func (t *Tile) Flip() {
	t.matrix.Flip()
}

// Rotates the tile 90 degrees clockwise
func (t *Tile) Rotate() {
	t.matrix.Rotate()
}

// Returns the state at point p
func (t *Tile) At(p Point) rune {
	return rune(t.matrix.At(p))
}

// Sets the state at point p
func (t *Tile) Set(p Point, v rune) {
	t.matrix.Set(p, int(v))
}

func (t *Tile) String() string {
	out := fmt.Sprintf("Tile %d:\n", t.ID)
	out += t.matrix.String()
	return out
}

func (t *Tile) checkBounds(p Point) {
	if p.X < 0 || p.X > t.Width() - 1 || p.Y < 0 || p.Y > t.Height() - 1 {
		panic(fmt.Sprintf("point out of bounds: %v", p))
	}
}

func (t *Tile) Width() int {
	return t.matrix.Width()
}

func (t *Tile) Height() int {
	return t.matrix.Height()
}

func (t *Tile) BorderTop() string {
	out := ""
	for i:=0; i<=t.Width()-1; i++ {
		out += string(t.At(Point{X: i}))
	}
	return out
}

func (t *Tile) BorderBottom() string {
	out := ""
	for i:=0; i<=t.Width()-1; i++ {
		out += string(t.At(Point{X: i, Y: t.Height()-1}))
	}
	return out
}

func (t *Tile) BorderLeft() string {
	out := ""
	for i:=0; i<=t.Width()-1; i++ {
		out += string(t.At(Point{Y: i}))
	}
	return out
}

func (t *Tile) BorderRight() string {
	out := ""
	for i:=0; i<=t.Width()-1; i++ {
		out += string(t.At(Point{X: t.Height()-1, Y: i}))
	}
	return out
}

// Stitch a list of tiles together into an array of strings, one line per row, after removing their borders
func Stitch(tiles []Tile) []string {
	s := make([]string, tiles[0].Height())
	for _, t := range tiles {
		for y := 1; y < tiles[0].Height()-1; y++ {
			s[y] = t.Column(y)[1:9]
		}
	}
	return s
}

func (t *Tile) Inner() string {
	s := ""
	for y:=1; y<t.Height(); y++ {
		for x:=1; x<=t.Width()-2;x++ {
			s += string(t.At(Point{X: x, Y: y}))
		}
	}
	return s
}

func (t *Tile) Column(row int) string {
	out := ""
	for i:=0; i <= t.Width()-1; i++ {
		out += string(t.At(Point{X: i, Y: row}))
	}
	return out
}

func (t *Tile) Variations() []Tile {
	var variations []Tile

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
