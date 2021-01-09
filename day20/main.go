package main

import (
	. "aoc"
	. "aoc/types"
	"math"
	"os"
)

type Board struct {
	tiles [][]*Tile
	dim int
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

func main() {
	Solve(Input("day20/input.txt"))
}

func Solve(input string) {
	tiles := parseInput(input)
	board := NewBoard(len(tiles))
	variations := tileVariations(tiles)
	EnsureUnique(variations)
	visited := NewIntSet()
	board.FillAndSolve(0, 0, variations, visited)
}

func parseInput(input string) []Tile {
	blocks := SplitByEmptyNewline(input)
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

// Fill the board by recursively checking every possible placement using pre-generated tile variations (8*N).
func (b *Board) FillAndSolve(row int, col int, allTiles []Tile, visited *IntSet) {
	// We've reached the end
	if row == b.dim {
		p1 := b.tiles[0][0].ID *
			b.tiles[0][b.dim-1].ID *
			b.tiles[b.dim-1][0].ID *
			b.tiles[b.dim-1][b.dim-1].ID

		//p2 := b.findMonsters()
		Printfln("P1: %d", p1)
		//Printfln("P2: %d", p2)

		// Todo: resolve backtracking efficiently without os.Exit()
		os.Exit(0)
	}
	for _, t := range allTiles {
		if visited.Contains(t.ID) {
			continue
		}
		// Check if we can place it below us
		if row > 0 && b.tiles[row-1][col].BorderBottom() != t.BorderTop() {
			continue
		}
		// Check if we can place it next to us
		if col > 0 && b.tiles[row][col-1].BorderRight() != t.BorderLeft() {
			continue
		}
		b.tiles[row][col] = &t

		visited.Add(t.ID)
		if col == b.dim-1 {
			b.FillAndSolve(row+1, 0, allTiles, visited)
		} else {
			b.FillAndSolve(row, col+1, allTiles, visited)
		}
		// We've hit a dead end, go back and try again
		visited.Remove(t.ID)
	}
}

func EnsureUnique(tiles []Tile) {
	seen := NewStringSet()
	for _, tt := range tiles {
		s := tt.String()
		if seen.Contains(s) {
			panic("Non-unique tile variation")
		}
		seen.Add(s)
	}
}