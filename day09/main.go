package main

import (
	. "aoc"
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	input := Ints(Input("day09/input.txt"))
	firstInvalid := Solve1(input, 25)
	p2 := Solve2(input, firstInvalid)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", firstInvalid)
	fmt.Println("p2:", p2)
}

func Solve1(input []int, preambleSize int) int {
	for i:=preambleSize; i < len(input); i++ {
		if !valid(input[i], input[i-preambleSize:i]) {
			return input[i]
		}
	}
	return 0
}

func valid(n int, preamble []int) bool {
	for i := 0; i < len(preamble); i++ {
		for j := 0; j < len(preamble); j++ {
			if i == j {
				continue
			}
			if n == preamble[i] + preamble[j] {
				return true
			}
		}
	}
	return false
}

func Solve2(input []int, n int) int {
	for size := 2; size < len(input); size++ {
		var set []int
		for i := 0; i < len(input); i++ {
			set = []int{}
			for j:=i; j < size; j++ {
				set = append(set, input[j])
			}
			if sum(set) == n {
				return min(set) + max(set)
			}
		}
	}

	return 0
}

func sum(set []int) int {
	n := 0
	for _, i := range set {
		n += i
	}
	return n
}

func min(n []int) int {
	sort.Ints(n)
	return n[0]
}

func max(n []int) int {
	sort.Ints(n)
	return n[len(n)-1]
}
