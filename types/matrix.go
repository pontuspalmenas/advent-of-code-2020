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

func (m *Matrix) Copy() *Matrix {
	d := make([][]int, m.Height())
	for i:=0; i< m.Height(); i++ {
		d[i] = make([]int, m.Width())
	}

	d = Copy2DIntSlice(m.data)
	return &Matrix{data: d}
}

func (m *Matrix) Flip() {
	for i:=0; i< m.Height(); i++ {
		for j:=0; j< m.Width()/2; j++ {
			c := m.data[i][j]
			m.data[i][j] = m.data[i][len(m.data[i])-1-j]
			m.data[i][len(m.data[i])-1-j] = c
		}
	}
}

// Rotates the Matrix 90 degrees clockwise
// todo: do in-place rotation, this is O(n^2) space
func (m *Matrix) Rotate() {
	if m.Width() != m.Height() {
		Panic("Rotate: unsupported NxM matrix")
	}

	n := m.Height()

	d := make([][]int, m.Height())
	for i:=0; i< m.Height(); i++ {
		d[i] = make([]int, m.Width())
	}

	for i:=0; i<n;i++ {
		for j:=0; j<n;j++ {
			d[i][j] = m.data[n - j - 1][i]
		}
	}

	m.data = Copy2DIntSlice(d)
}

// Returns the state at point p
func (m *Matrix) At(p Point) int {
	m.checkBounds(p)
	return m.data[p.Y][p.X]
}

// Sets the state at point p
func (m *Matrix) Set(p Point, v int) {
	m.checkBounds(p)
	m.data[p.Y][p.X] = v
}

// Returns a printable Matrix
// todo: print ints as numbers, not their rune. but that breaks Tile.String()...
func (m *Matrix) String() string {
	out := ""
	for y := 0; y <= m.Height()-1; y++ {
		for x := 0; x <= m.Width()-1; x++ {
			out += string(rune(m.data[y][x]))
		}
		out += "\n"
	}
	return out
}

func (m *Matrix) ToStringSlice() []string {
	var out []string
	for y := 0; y <= m.Height()-1; y++ {
		line := ""
		for x := 0; x <= m.Width()-1; x++ {
			line += string(rune(m.data[y][x]))
		}
		out = append(out, line)
	}
	return out
}

func (m *Matrix) checkBounds(p Point) {
	if p.X < 0 || p.X > m.Width() - 1 || p.Y < 0 || p.Y > m.Height() - 1 {
		Panic("point out of bounds: %v", p)
	}
}

func (m *Matrix) Width() int {
	return len(m.data[0])
}

func (m *Matrix) Height() int {
	return len(m.data)
}
