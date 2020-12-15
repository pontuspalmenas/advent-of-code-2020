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
	p := &Point{}
	dir := 90 // E
	for _, s := range input {
		var action rune
		var v int
		Sscanf(s, "%c%d", &action, &v)
		switch action {
		case 'N':
			north(p, v)
		case 'S':
			south(p, v)
		case 'E':
			east(p, v)
		case 'W':
			west(p, v)
		case 'L':
			dir = dir - v
			if dir < 0 {
				dir = 360 + dir
			}
		case 'R':
			dir = (dir + v) % 360
			if dir < 0 {
				dir = 360 + dir
			}
		case 'F':
			switch dir {
			case 0:
				north(p, v)
			case 90:
				east(p, v)
			case 180:
				south(p, v)
			case 270:
				west(p, v)
			default:
				panic("unhandled position")
			}
		}
	}

	return Manhattan(Point{}, *p)
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
	p := &Point{}
	wp := &Point{X: 10, Y: 1}
	for _, s := range input {
		var action rune
		var v int
		Sscanf(s, "%c%d", &action, &v)
		switch action {
		case 'N':
			north(wp, v)
		case 'S':
			south(wp, v)
		case 'E':
			east(wp, v)
		case 'W':
			west(wp, v)
		case 'L':
			for i := 0; i < v/90; i++ {
				wp = &Point{X: -wp.Y, Y: wp.X}
			}
		case 'R':
			for i := 0; i < v/90; i++ {
				wp = &Point{X: wp.Y, Y: -wp.X}
			}
		case 'F':
			for i := 0; i < v; i++ {
				p = &Point{X: p.X + wp.X, Y: p.Y + wp.Y}
			}
		}
	}

	return Manhattan(Point{}, *p)
}
