package main

import (
	. "aoc"
)

// Tile is a 10x10 grid of characters
type Tile struct {
	ID int
	Data [][]rune
}

func NewTileFromString(s string) Tile {
	d := make([][]rune, 10)
	for i:=0; i<10; i++ {
		d[i] = make([]rune, 10)
	}

	lines := Lines(s)
	id := Int(Regex(`Tile (\d+)`, lines[0])[0])
	tile := Tile{ID: id, Data: d}
	lines = lines[1:]
	for i := 0; i <= 10-1; i++ {
		for j := 0; j <= 10-1; j++ {
			tile.Data[i][j] = rune(lines[i][j])
		}
	}

	return tile
}

func (t *Tile) Rotate() {
	t.Data = Copy2DRuneSlice(Rotate(t.Data))
}

func (t *Tile) Flip() {
	t.Data = Copy2DRuneSlice(Flip(t.Data))
}

func Flip(data [][]rune) [][]rune {
	h := len(data)
	w := len(data[0])

	for i:=0; i< h; i++ {
		for j:=0; j<w/2; j++ {
			c := data[i][j]
			data[i][j] = data[i][len(data[i])-1-j]
			data[i][len(data[i])-1-j] = c
		}
	}

	return data
}

func Rotate(data [][]rune) [][]rune {
	h := len(data)
	w := len(data[0])

	d := make([][]rune, h)
	for i:=0; i<h; i++ {
		d[i] = make([]rune, w)
	}

	for i:=0; i<h;i++ {
		for j:=0; j<h;j++ {
			d[i][j] = data[h - j - 1][i]
		}
	}

	return d
}

func (t *Tile) Copy() *Tile {
	d := make([][]rune, 10)
	for i:=0; i< 10; i++ {
		d[i] = make([]rune, 10)
	}

	d = Copy2DRuneSlice(t.Data)
	return &Tile{ID: t.ID, Data: d}
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

func (t *Tile) String() string {
	out := ""
	for y := 0; y <= 10-1; y++ {
		for x := 0; x <= 10-1; x++ {
			out += string(t.Data[y][x])
		}
		out += "\n"
	}
	return out
}

func (t *Tile) BorderTop() string {
	out := ""
	for i:=0; i<=10-1; i++ {
		out += string(t.Data[0][i])
	}
	return out
}

func (t *Tile) BorderBottom() string {
	out := ""
	for i:=0; i<=10-1; i++ {
		out += string(t.Data[9][i])
	}
	return out
}

func (t *Tile) BorderLeft() string {
	out := ""
	for i:=0; i<=10-1; i++ {
		out += string(t.Data[i][0])
	}
	return out
}

func (t *Tile) BorderRight() string {
	out := ""
	for i:=0; i<=10-1; i++ {
		out += string(t.Data[i][9])
	}
	return out
}

