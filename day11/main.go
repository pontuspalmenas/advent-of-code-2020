package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day11/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []string) int {
	grid := parseGrid(input)

	for {
		var changes int
		grid, changes = next(grid)
		if changes == 0 {
			return countOccupied(grid)
		}
	}
}

func parseGrid(input []string) [][]rune {
	grid := [][]rune{}
	for _, line := range input {
		row := []rune{}
		for _, col := range line {
			row = append(row, col)
		}
		grid = append(grid, row)
	}
	return grid
}

func next(grid [][]rune) ([][]rune, int) {
	changes := 0
	temp := copyGrid(grid)
	for i, row := range grid {
		for j, col := range row {
			if col == 'L' && occupiedAdjacents(i,j,grid) == 0 {
				temp[i][j] = '#'
				changes++
			}
			if col == '#' {
				if occupiedAdjacents(i, j, grid) >= 4 {
					temp[i][j] = 'L'
					changes++
				}
			}
		}
	}
	return temp, changes
}

func copyGrid(grid [][]rune) [][]rune {
	g := make([][]rune, len(grid))
	for i := range grid {
		g[i] = make([]rune, len(grid[i]))
		copy(g[i], grid[i])
	}
	return g
}

func occupiedAdjacents(i, j int, grid [][]rune) int {
	occ := 0
	for y := i-1; y <i+2; y++ {
		for x := j-1; x <j+2; x++ {
			if !(i == y && j == x) && occupied(y, x, grid) {
				occ++
			}
		}
	}

	return occ
}

func countOccupied(grid [][]rune) int {
	count := 0
	for i, row := range grid {
		for j := range row {
			if occupied(i,j, grid) {
				count++
			}
		}
	}
	return count
}

func occupied(i, j int, grid [][]rune) bool {
	if i > len(grid)-1 || i < 0 { return false }
	if j > len(grid[i])-1 || j < 0 { return false }
	return grid[i][j] == '#'
}

func draw(grid [][]rune) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func Solve2(input []string) int {
	return 0
}
