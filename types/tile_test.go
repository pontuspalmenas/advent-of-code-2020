package types

import (
	. "aoc"
	"testing"
)

func TestNewTile(t *testing.T) {
	expected :=
`Tile 2:
0000
1111
2222
3333
`

	tile := NewTileFromString(expected)
	AssertEq(t, tile.String(), expected)
}

func TestTile_Rotate(t *testing.T) {
	input :=
`Tile 3:
1234
5678
9012
3456
`
	expected :=
`Tile 3:
3951
4062
5173
6284
`

	tile := NewTileFromString(input)
	tile.Rotate()
	AssertEq(t, tile.String(), expected)
}

func TestTile_Borders(t *testing.T) {
	input :=
`Tile 4:
1234
5678
9012
3456
`

	left := "1593"
	right := "4826"
	top := "1234"
	bottom := "3456"

	tile := NewTileFromString(input)
	AssertEq(t, tile.BorderLeft(), left)
	AssertEq(t, tile.BorderRight(), right)
	AssertEq(t, tile.BorderTop(), top)
	AssertEq(t, tile.BorderBottom(), bottom)

}