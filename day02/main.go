package main

import (
	. "aoc"
	"fmt"
	"strings"
)

func main() {
	input := Lines(Input("day02/input.txt"))
	fmt.Println(Solve1(input))
	fmt.Println(Solve2(input))
}

func Solve1(ss []string) int {
	valid := 0
	for _, s := range ss {
		// todo: replace with regex
		minmax := Ints(strings.ReplaceAll(s, "-", " "))
		target := strings.Split(s, " ")[1][0]

		// todo: replace with map
		count := 0
		chars := strings.Split(s, ":")[1][1:]
		for i:=0; i<len(chars); i++ {
			if chars[i] == target {
				count++
			}
		}
		if count >= minmax[0] && count <= minmax[1] {
			valid++
		}
		fmt.Printf("%d-%d %d: %s\n", minmax[0], minmax[1], target, chars)
	}

	return valid
}

func Solve2(ss []string) int {
	return 0
}
