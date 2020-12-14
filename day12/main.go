package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day12/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []string) int {
	position := &Point{0,0}
	dir := 90 // E
	for _, s := range input {
		var action rune
		var size int
		Sscanf(s, "%c%d", &action, &size)
		switch action {
		case 'N':
			north(position, size)
		case 'S':
			south(position, size)
		case 'E':
			east(position, size)
		case 'W':
			west(position, size)
		case 'L':
			dir = dir - size
			if dir < 0 {
				dir = 360 + dir
			}
		case 'R':
			dir = (dir + size) % 360
			if dir < 0 {
				dir = 360 + dir
			}
		case 'F':
			switch dir {
			case 0:
				north(position, size)
			case 90:
				east(position, size)
			case 180:
				south(position, size)
			case 270:
				west(position, size)
			default:
				panic("unhandled position")
			}
		}
	}

	return Manhattan(Point{0,0}, *position)
}

func north(p *Point, distance int) {
	p.Y += distance
}

func south(p *Point, distance int) {
	p.Y -= distance
}

func east(p *Point, distance int) {
	p.X += distance
}

func west(p *Point, distance int) {
	p.X -= distance
}

func Solve2(input []string) int {
	return 0
}
