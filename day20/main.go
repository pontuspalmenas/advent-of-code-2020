package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
	"math"
	"os"
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

// Fill the board by recursively checking every possible placement using pre-generated tile variations (8*N).
func (b *Board) Fill(row int, col int, allTiles []Tile, visited *IntSet) {
	// We've reached the end
	if row == b.dim {
		Printfln("P1: %d", b.tiles[0][0].ID *
			b.tiles[0][b.dim-1].ID *
			b.tiles[b.dim-1][0].ID *
			b.tiles[b.dim-1][b.dim-1].ID)
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
			b.Fill(row+1, 0, allTiles, visited)
		} else {
			b.Fill(row, col+1, allTiles, visited)
		}
		// We've hit a dead end, go back and try again
		visited.Remove(t.ID)
	}
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
	visited := NewIntSet()
	board.Fill(0, 0, variations, visited)

	//board.Print()

	return board.tiles[0][0].ID *
		board.tiles[0][board.dim-1].ID *
		board.tiles[board.dim-1][0].ID *
		board.tiles[board.dim-1][board.dim-1].ID
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
