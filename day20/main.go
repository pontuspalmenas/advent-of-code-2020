package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	input := SplitByEmptyNewline(Input("day20/input.txt"))
	Solve(input)
	fmt.Println(time.Since(start))
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
func (b *Board) FillAndSolve(row int, col int, allTiles []Tile, visited *IntSet) {
	// We've reached the end
	if row == b.dim {
		p1 := b.tiles[0][0].ID *
			b.tiles[0][b.dim-1].ID *
			b.tiles[b.dim-1][0].ID *
			b.tiles[b.dim-1][b.dim-1].ID

		p2 := b.findMonsters()
		Printfln("P1: %d", p1)
		Printfln("P2: %d", p2)

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


// Finds sea monsters and returns "roughness" (number of #'s not part of sea monster)
func (b *Board) findMonsters() int {
	monster := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	

	// Remove frames around tiles, and stitch them all together
	sea := make([]string, 0)
	for row := 0; row <= 9*b.dim; row++ {
		line := ""
		for _, c := range b.tiles[row/10] {
			if row % 10 == 0 { // Skip top and bottom borders
				continue
			}
			line += c.Column(row%c.Height())[1:9] // Skip left and right borders
		}
		if line != "" {
			sea = append(sea, line)
		}
	}

	for _, s := range sea {
		fmt.Println(s)
	}

	Printfln("%d*%d", len(sea), len(sea[0]))

	// todo: compile error
	sort.Strings(monster)

	return 0
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

func Solve(input []string) {
	tiles := parseInput(input)
	board := NewBoard(len(tiles))
	variations := tileVariations(tiles)
	visited := NewIntSet()
	board.FillAndSolve(0, 0, variations, visited)
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
