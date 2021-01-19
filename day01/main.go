package main

import (
	. "aoc"
	"fmt"
)

func main() {
	input := Ints(Input("day01/input.txt"))
	fmt.Println(Solve1(input))
	fmt.Println(Solve2(input))
}

// Todo: unoptimized. fast enough though. use a lookup table/cache to improve.

func Solve1(input []int) int {
	for _, i := range input {
		for _, j := range input {
			if i+j == 2020 {
				return i * j
			}
		}
	}
	return 0
}

func Solve2(input []int) int {
	for _, i := range input {
		for _, j := range input {
			for _, k := range input {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}
	return 0
}
