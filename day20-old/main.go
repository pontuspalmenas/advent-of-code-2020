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
	tiles [][]*Tile_OLD
	dim int
}

// Creates a Board with "NxN=size" dimensions
func NewBoard(size int) *Board {
	dim := int(math.Sqrt(float64(size)))
	tiles := make([][]*Tile_OLD, dim)
	for i:=0; i<dim; i++ {
		tiles[i] = make([]*Tile_OLD, dim)
	}
	return &Board{tiles: tiles, dim: dim}
}

// Fill the board by recursively checking every possible placement using pre-generated tile variations (8*N).
func (b *Board) FillAndSolve(row int, col int, allTiles []Tile_OLD, visited *IntSet) {
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

var monster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

// Finds sea monsters and returns "roughness" (number of #'s not part of sea monster)
func (b *Board) findMonsters() int {
	// Remove frames around tiles, and stitch them all together
	sea := make([]string, 0)
	for _, r := range b.tiles {
		sea = append(sea, b.Stitch(r)...)
	}

	m := NewMatrix(96,96)
	for y,r := range sea {
		for x,c := range r {
			m.Set(Point{X:x, Y: y}, int(c))
		}
	}

	for i:=0; i<2; i++ {
		FindMonsters(m.ToStringSlice())
		for j := 0; j < 3; j++ {
			m.Rotate()
			FindMonsters(m.ToStringSlice())
		}
		m.Flip()
	}

	Printfln("%d*%d", len(sea), len(sea[0]))

	// todo: compile error
	sort.Strings(monster)

	return 0
}

func FindMonsters(sea []string) int {
	h := len(monster)
	w := len(monster[0])

	spots := NewPointSet()
	for row := 0; row < len(sea)-h-1; row++ {
		for col:=0; col < len(sea[0])-w-1; col++ {
			area := []string{sea[row][col:col+w],sea[row+1][col:col+w],sea[row+1][col:col+w]}
			if FindMonster(area) {
				for r := 0; r < h; r++ {
					for c := 0; c < w; c++ {
						if monster[r][c] == '#' {
							spots.Add(Point{X: col+c, Y: row+r})
						}
					}
				}
				println("MONSTER MONSTER MONSTER")
				os.Exit(0)
			}
		}
	}

	return 0
}

func FindMonster(area []string) bool {
	found := true
	for y:=0; y < len(monster); y++ {
		for x := range area[y] {
			if monster[y][x] == '#' && area[y][x] == '.' {
				found = false
			}
		}
	}
	return found
}

// Stitch a list of tiles together into an array of strings, one line per row, after removing their borders
func (b *Board) Stitch(tiles []*Tile_OLD) []string {
	s := make([]string, tiles[0].Height()-2)
	for _, t := range tiles {
		for y := 0; y < tiles[0].Height()-2; y++ {
			s[y] += t.Column(y+1)[1:t.Width()-1]
		}
	}
	return s
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

func parseInput(blocks []string) []Tile_OLD {
	var tiles []Tile_OLD
	for _, b := range blocks {
		tiles = append(tiles, NewTileFromString_OLD(b))
	}
	return tiles
}

func tileVariations(tiles []Tile_OLD) []Tile_OLD {
	var variations []Tile_OLD
	for _, t := range tiles {
		variations = append(variations, t.Variations()...)
	}
	return variations
}
