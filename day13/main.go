package main

import (
	. "aoc"
	"aoc/types"
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

	return (when-myDeparture)*earliestBus
}

type Bus struct {
	id int
	offset int
}

func Solve2(input []string) int {
	busses := getBussesWithOffset(input[1])

	t := 0
	now := 1
	used := types.NewIntSet()
	found := false
	for !found {
		t += now
		found = true
		for _, bus := range busses {
			if (t + bus.offset) % bus.id != 0 {
				found = false
				break
			}
			if !used.Contains(bus.id) {
				used.Add(bus.id)
				now = bus.id * now
			}
		}
	}

	return t
}

func getBussesWithOffset(in string) (busses []Bus) {
	for i, s := range SplitByComma(in) {
		if s != "x" {
			busses = append(busses, Bus{Int(s), i})
		}
	}
	return busses
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
