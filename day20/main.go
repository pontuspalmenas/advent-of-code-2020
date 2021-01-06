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
		tiles = append(tiles, NewTileFromString(b))
	}
	return tiles
}
