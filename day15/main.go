package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	p1 := Solve1("0,13,16,17,1,10,6")
	p2 := Solve2("0,13,16,17,1,10,6")
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input string) int {
	spoken := initSpoken()
	for i, n := range Ints(input) {
		spoken[i] = n
	}
	for i := len(Ints(input)); i < 10; i++ {
		Printfln("Turn %d", i+1)
		prev := spoken[i-1]
		w := when(spoken, prev, i)
		Printfln("prev: %d, when: %d", prev, w)
		if w == -1 {
			spoken[i] = 0
		} else {
			spoken[i] = (i-1)-(w-1)
			Printfln("%d - %d = %d", i-1, w-1, spoken[i])
		}
	}

	return spoken[2019]
}

func initSpoken() []int {
	n := make([]int, 2020)
	for i:=0; i<len(n); i++{
		n[i] = -1
	}
	return n
}

func when(spoken []int, n int, turn int) int {
	for i := turn-2; i>0; i-- {
		if spoken[i] == n {
			return i
		}
	}

	return -1
}

func Solve2(input string) int {
	return 0
}
