package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := "0,13,16,17,1,10,6"
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input string) int {
	return play(input, 2020)
}

func Solve2(input string) int {
	return play(input, 30000000)
}

func play(input string, turns int) int {
	spoken := make(map[int]int)
	init := Ints(input)
	prev := -1
	for i, n := range init {
		spoken[n] = i
		prev = n
	}
	for i := len(init); i < turns; i++ {
		var val int
		prevPos, seen := spoken[prev]
		if seen {
			val = (i - 1) - prevPos
		} else {
			val = 0
		}
		if prev >= 0 {
			spoken[prev] = i - 1
		}
		prev = val
	}

	return prev
}
