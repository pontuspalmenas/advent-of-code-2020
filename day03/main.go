package main

import (
	. "aoc"
	"fmt"
)

func main() {
	input := Lines(Input("day03/input.txt"))
	fmt.Println(Solve1(input))
	fmt.Println(Solve2(input))
}

func Solve1(ss []string) int {
	return countTrees(ss, 3, 1)
}

func countTrees(ss []string, right int, down int) int {
	trees := 0
	y := 0
	x := 0
	width := len(ss[0])
	lines := len(ss)
	for {
		y = y + down
		x = (x + right) % width

		if y >= lines {
			break
		}

		if ChrAt(ss[y], x) == "#" {
			trees++
		}
		//draw(ss[y], x)
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
