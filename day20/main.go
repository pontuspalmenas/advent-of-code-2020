package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	input := SplitByEmptyNewline(Input("day20/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

type Board struct {
	tiles [][]*Tile

	// can be calculated as we go, but this makes the code cleaner
	dim int
	usedH int
	usedW int
}

// Creates a Board with "NxN=size" dimensions
func NewBoard(size int) *Board {
	dim := int(math.Sqrt(float64(size)))
	tiles := make([][]*Tile, dim)
	for i:=0; i<dim; i++ {
		tiles[i] = make([]*Tile, dim)
	}
	return &Board{tiles: tiles, dim: dim}
}

func (b *Board) Place(t Tile) {
	b.tiles[b.usedH][b.usedW] = &t
	b.usedW++
	if b.usedW > len(b.tiles[0])-1 {
		b.usedW = 0
		b.usedH++
	}
}

func (b *Board) AllLinedUp() bool {
	return false
}

func (b *Board) LinedUp(side string, t1 Tile, t2 Tile) bool {
	switch side {
	case "left":
	case "right":
	case "top":
	case "bottom":
	}

	return false
}

func (b *Board) Print() {
	for _, r := range b.tiles {
		for line := 0; line < b.tiles[0][0].Height(); line++ {
			for _, c := range r {
				if c != nil {
					fmt.Print(c.Column(line) + "  ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func Solve1(input []string) int {
	tiles := parseInput(input)
	board := NewBoard(len(tiles))
	variations := tileVariations(tiles)
	Printfln("variations: %d", len(variations))
	for _, t := range tiles {
		board.Place(t)
	}
	//board.Print()

	return 0
}

func Solve2(input []string) int {
	return 0
}

func parseInput(blocks []string) []Tile {
	var tiles []Tile
	for _, b := range blocks {
		tiles = append(tiles, NewTileFromString(b))
	}
	return tiles
}

func tileVariations(tiles []Tile) []Tile {
	var variations []Tile
	for _, t := range tiles {
		variations = append(variations, t.Variations()...)
	}
	return variations
}
