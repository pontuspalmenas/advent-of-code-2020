package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day13/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []string) int {
	myDeparture := Int(input[0])
	busses := Ints(input[1])

	earliestBus, when := findBus(myDeparture, busses)

	printSchedule(busses)

	return (when-myDeparture)*earliestBus
}

func Solve2(input []string) int {
	return 0
}

func findBus(myDeparture int, busses []int) (int, int) {
	i := myDeparture
	for {
		for _, bus := range busses {
			if i % bus == 0 {
				return bus, i
			}
		}

		i++
	}
}

func printSchedule(busses []int) {
	fmt.Print("time\t")
	for _, bus := range busses {
		fmt.Printf("bus %d\t", bus)
	}
	fmt.Println()

	for i:=929; i<950; i++ {
		fmt.Printf("%d\t", i)
		for _, bus := range busses {
			if i % bus == 0 {
				fmt.Printf("\t  D\t")
			} else {
				fmt.Printf("\t  .\t")
			}
		}
		fmt.Println()
	}
}
