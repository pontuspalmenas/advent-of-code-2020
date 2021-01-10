package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
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
	visited := NewIntSet()
	board.FillAndSolve(0, 0, variations, visited)
}

// Takes a set up board and solves for part 2
func Solve2(b *Board) int {
	sea := b.Stitch()
	//PrintSea(sea)

	// Rotate, Flip, until we find the monsters
	for i:=0; i<2; i++ {
		for j := 0; j < 3; j++ {
			if Roughness(sea) > 0 {
				panic("solved")
			}
			sea = Rotate(sea)
		}
		sea = Flip(sea)
	}

	return 0
}

// Remove frames around tiles and stitch together
func (b *Board) Stitch() [][]rune {
	dim := b.dim * 8 // Tiles are 10x10, remove frame => 8
	s := Make2dRuneSlice(dim, dim)

	i:=0
	j:=0
	for y:=0;y<96;y++ {
		for x:=0;x<96;x++ {
			//Printfln("s[%d][%d] = b.tiles[%d][%d].Data[%d][%d]", y,x,y/8,x/8,i+1,j+1)
			s[y][x] = b.tiles[y/8][x/8].Data[i+1][j+1]
			j = (j+1)%8
		}
		i = (i+1)%8
	}

	return s
}

func Roughness(sea [][]rune) int {
	h := len(monster)
	w := len(monster[0])

	for row := 0; row < len(sea)-h-1; row++ {
		for col:=0; col < len(sea[0])-w-1; col++ {
			area := [][]rune{
				sea[row][col:col+w],
				sea[row+1][col:col+w],
				sea[row+1][col:col+w]}
			if FindMonster(area) {
				for r := 0; r < h; r++ {
					for c := 0; c < w; c++ {
						if monster[r][c] == '#' {
							sea[row+r][col+c] = 'O'
						}
					}
				}
				PrintSea(sea)
				println("MONSTER MONSTER MONSTER")
				os.Exit(0)
			}
		}
	}

	return 0
}

var monster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

func FindMonster(area [][]rune) bool {
	found := true
	for y:=0; y < len(monster); y++ {
		for x:=0; x < len(monster[0]); x++ {
			if monster[y][x] == '#' && area[y][x] == '.' {
				found = false
			}
		}
	}
	return found
}

func PrintSea(sea [][]rune) {
	for i:=0;i<96;i++ {
		for j:=0;j<96;j++ {
			fmt.Printf("%c", sea[i][j])
		}
		fmt.Println()
	}
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

		p2 := Solve2(b)
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
