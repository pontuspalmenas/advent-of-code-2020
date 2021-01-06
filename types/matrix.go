package types

import (
	. "aoc"
)

type Matrix struct {
	data [][]int
}

func NewMatrix(width, height int) Matrix {
	d := make([][]int, height)
	for i:=0; i<height; i++ {
		d[i] = make([]int, width)
	}

	return Matrix{data: d}
}

func (t *Matrix) Flip() {

}

// Rotates the Matrix 90 degrees clockwise
func (t *Matrix) Rotate() {
	if t.Width() != t.Height() {
		Panic("Rotate: unsupported NxM matrix")
	}

	n := t.Height()

	d := make([][]int, t.Height())
	for i:=0; i<t.Height(); i++ {
		d[i] = make([]int, t.Width())
	}

	for i:=0; i<n;i++ {
		for j:=0; j<n;j++ {
			d[i][j] = t.data[n - j - 1][i]
		}
	}

	t.data = Copy2DIntSlice(d)
}

// Returns the state at point p
func (t *Matrix) At(p Point) int {
	t.checkBounds(p)
	return t.data[p.Y][p.X]
}

// Sets the state at point p
func (t *Matrix) Set(p Point, v int) {
	t.checkBounds(p)
	t.data[p.Y][p.X] = v
}

// Returns a printable Matrix
// todo: print ints as numbers, not their rune. but that breaks Tile.String()...
func (t *Matrix) String() string {
	out := ""
	for y := 0; y <= t.Height()-1; y++ {
		for x := 0; x <= t.Width()-1; x++ {
			out += string(rune(t.data[y][x]))
		}
		out += "\n"
	}
	return out
}

func (t *Matrix) checkBounds(p Point) {
	if p.X < 0 || p.X > t.Width() - 1 || p.Y < 0 || p.Y > t.Height() - 1 {
		Panic("point out of bounds: %v", p)
	}
}

func (t *Matrix) Width() int {
	return len(t.data[0])
}

func (t *Matrix) Height() int {
	return len(t.data)
}