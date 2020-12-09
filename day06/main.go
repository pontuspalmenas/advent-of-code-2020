package main

import (
	. "aoc"
	"fmt"
	"strings"
)

func main() {
	input := Input("day06/input.txt")
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input string) int {
	sum := 0
	for _, group := range SplitByEmptyNewline(input) {
		answers := make(map[rune]bool)
		for _, s := range strings.Split(group, "\n") {
			for _, r := range s {
				answers[r] = true
			}
		}
		for _, answer := range answers {
			if answer {
				sum++
			}
		}
	}

	return sum
}

func Solve2(input string) int {
	sum := 0
	for _, group := range SplitByEmptyNewline(input) {
		answers := make(map[rune]int)
		for _, s := range strings.Split(group, "\n") {
			for _, r := range s {
				answers[r] = answers[r] + 1
			}
		}
		for _, answer := range answers {
			if answer == len(strings.Split(group, "\n")) {
				sum++
			}
		}
	}
	return sum
}
