package main

import (
	. "aoc"
	"fmt"
)

func main() {
	input := Lines(Input("day03/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(ss []string) int {
	return countTrees(ss, 3, 1)
}

func countTrees(ss []string, right int, down int) int {
	trees := 0
	x := 0
	for y := down; y < len(ss); y += down {
		draw(ss[y], x)
		x = (x + right) % len(ss[0]) // map is wrapped

		if ss[y][x] == '#' {
			trees++
		}
	}

	return trees
}

func Solve2(ss []string) int {
	return countTrees(ss, 1, 1) *
	countTrees(ss, 3, 1) *
	countTrees(ss, 5, 1) *
	countTrees(ss, 7, 1) *
	countTrees(ss, 1, 2)
}

func draw(s string, x int) {
	l := []rune(s)
	if l[x] == '#' {
		l[x] = 'X'
	} else {
		l[x] = 'O'
	}

	fmt.Println(string(l))
}
