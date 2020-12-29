package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
	"regexp"
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
	return tiles.Size()
}

func Solve2(input []string) int {
	tiles := load(input)

	for i := 1; i <= 100; i++ {
		tiles = next(tiles)
	}

	return tiles.Size()
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
func load(input []string) *PointSet {
	blacks := NewPointSet() // no need to keep unmarked/white tiles
	for _, row := range input {
		moves := split(row)
		p := Point{}
		for _, m := range moves {
			dir := directions[m]
			p.X += dir.dx
			p.Y	+= dir.dy
		}
		if blacks.Contains(p) {
			blacks.Remove(p)
		} else {
			blacks.Add(p)
		}
	}
	return blacks
}

func neighbors(p Point) *PointSet {
	ns := NewPointSet()
	for _, dir := range directions {
		ns.Add(Point{X: p.X + dir.dx, Y: p.Y + dir.dy})
	}
	return ns
}

func next(blacks *PointSet) *PointSet {
	out := NewPointSet()
	for _, p := range blacks.ToSlice() {
		count := blackNeighbors(blacks, p)
		if count == 1 || count == 2 {
			out.Add(p)
		}
		for _, np := range neighbors(p).ToSlice() {
			count = blackNeighbors(blacks, np)
			if count == 2 {
				out.Add(np)
			}
		}
	}
	return out
}

// count black tiles neighboring my Point p
func blackNeighbors(blacks *PointSet, p Point) int {
	count := 0
	for _, np := range neighbors(p).ToSlice() {
		if blacks.Contains(np) {
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


