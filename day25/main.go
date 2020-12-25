package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day25/input.txt"))
	p1 := Solve1(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
}

func Solve1(input []string) int {
	cardPubKey := Int(input[0])
	doorPubKey := Int(input[1])

	var cardLoopSize int
	value := 1
	for i := 1; ; i++ {
		value *= 7
		value %= 20201227
		if value == cardPubKey {
			cardLoopSize = i
			break
		}
	}

	value = 1
	for i := 1; i <= cardLoopSize; i++ {
		value *= doorPubKey
		value %= 20201227
	}

	return value
}

