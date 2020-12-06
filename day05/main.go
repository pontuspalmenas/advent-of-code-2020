package main

import (
	. "aoc"
	"fmt"
)

func main() {
	input := Lines(Input("day05/input.txt"))
	fmt.Println(Solve1(input))
	fmt.Println(Solve2(input))
}

func Solve1(input []string) int {
	for _, s := range input {
		r := row(s[:7])
		c := col(s[7:])
		fmt.Printf("%s: row %d, column %d, seat ID %d\n", s, r, c, r*8+c)
	}
	return 0
}

func row(s string) int {
	l := 0
	h := 127
	for i := 0; i < 7; i++ {
		if s[i] == 'F' {
			h = h / 2
			fmt.Printf("F means to take the lower half, keeping rows %d through %d\n", l, h)
		} else {
			l = h / 2
			fmt.Printf("B means to take the upper half, keeping rows %d through %d\n", l, h)
		}
	}
	return l
}

func col(s string) int {
	return 0
}

func Solve2(input []string) int {
	return 0
}
