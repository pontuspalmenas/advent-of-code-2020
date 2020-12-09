package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := Ints(Input("dayxx/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []int) int {
	return 0
}

func Solve2(input []int) int {
	return 0
}
