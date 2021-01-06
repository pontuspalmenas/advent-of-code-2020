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
			tile.Set(Point{j, i}, rune(lines[i][j]))
		}
	}
	return tile
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

