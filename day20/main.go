package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
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

const TileSize = 10

func Solve1(input []string) int {
	tiles := parseInput(input)
	for _, t := range tiles {
		Printfln("%v", t.String())
	}
	return 0
}

func Solve2(input []string) int {
	return 0
}

func parseInput(blocks []string) []Tile {
	var tiles []Tile
	for _, b := range blocks {
		lines := Lines(b)
		tile := NewTile(Int(Regex(`Tile (\d+)`, lines[0])[0]), TileSize, TileSize)
		lines = lines[1:]
		for i := 0; i <= TileSize-1; i++ {
			for j := 0; j <= TileSize-1; j++ {
				tile.Set(Point{j, i}, rune(lines[i][j]))
			}
		}
		tiles = append(tiles, tile)
	}
	return tiles
}
