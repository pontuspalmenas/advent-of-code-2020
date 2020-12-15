package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := Ints(Input("day10/input.txt")) // sort it first: $ sort -nr input.txt -o input.txt
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []int) int {
	input = append(input, 0) // add the outlet (0 jolts)
	adaptors := []int{input[0] + 3}
	adaptors = append(adaptors, input...)

	diffsOne := 0
	diffsThree := 0
	for i:=0; i<len(adaptors)-1; i++ {
		next := adaptors[i+1]
		this := adaptors[i]
		if this - next == 1 {
			diffsOne++
		} else if this - next == 3 {
			diffsThree++
		} else {
			panic(this)
		}
	}

	return diffsOne * diffsThree
}

func Solve2(input []int) int {
	input = append(input, 0) // add the outlet (0 jolts)
	adaptors := []int{input[0] + 3}
	adaptors = append(adaptors, input...)



	return 0
}
