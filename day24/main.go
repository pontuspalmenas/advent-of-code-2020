package main

import (
	. "aoc"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day24/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []string) int {
	tiles := load(input)
	return len(tiles)
}

func Solve2(input []string) int {
	tiles := load(input)

	for i := 1; i <= 100; i++ {
		tiles = next(tiles)
		Printfln("Day %d: %d", i, len(tiles))
	}

	return 0
}

type Direction struct {
	dy,dx int
}
var directions = map[string]Direction{
	"e":{dx: 1},
	"w":{dx: -1},
	"se":{dy: 1},
	"sw":{dx: -1, dy: 1},
	"ne":{dx: 1, dy: -1},
	"nw":{dy: -1},
}

// we represent the hexagon grid using axial coordinates:
// https://www.redblobgames.com/grids/hexagons/#coordinates
func load(input []string) []Point {
	var blacks []Point // no need to keep unmarked/white tiles
	for _, row := range input {
		moves := split(row)
		p := Point{}
		for _, m := range moves {
			dir := directions[m]
			p.X += dir.dx
			p.Y	+= dir.dy
		}
		if contains(blacks, p) {
			blacks = remove(blacks, p)
		} else {
			blacks = append(blacks, p)
		}

	}
	return blacks
}

func contains(blacks []Point, p Point) bool {
	for _, b := range blacks {
		if b == p {
			return true
		}
	}
	return false
}

func remove(blacks []Point, p Point) []Point {
	var out []Point
	for _, b := range blacks {
		if b != p {
			out = append(out, b)
		}
	}
	return out
}

func neighbors(p Point) []Point {
	var out []Point
	for _, s := range strings.Split("se|sw|ne|nw|e|w", "|") {
		dir := directions[s]
		out = append(out, Point{X: p.X + dir.dx, Y: p.Y + dir.dy})
	}
	return out
}

func next(blacks []Point) []Point {
	var out []Point
	for _, p := range blacks {
		count := blackNeighbors(blacks, p)
		if count == 1 || count == 2 {
			out = append(out, p)
		}
		for _, np := range neighbors(p) {
			count = blackNeighbors(blacks, np)
			if count == 2 {
				out = append(out, np)
			}
		}
	}
	return out
}

// count black tiles neighboring my Point p
func blackNeighbors(blacks []Point, p Point) int {
	count := 0
	for _, np := range neighbors(p) {
		if contains(blacks, np) {
			count++
		}
	}
	return count
}

func split(s string) []string {
	out := []string{}
	r := regexp.MustCompile(`se|sw|ne|nw|e|w`)
	out = r.FindAllString(s, -1)
	return out
}


