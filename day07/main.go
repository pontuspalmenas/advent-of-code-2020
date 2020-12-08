package main

import (
	. "aoc"
	"fmt"
	"strings"
)

func main() {
	input := Lines(Input("day07/input.txt"))
	fmt.Println(Solve1(input))
	fmt.Println(Solve2(input))
}

type bag struct {
	
}

var bags map[string]int

func Solve1(input []string) int {
	bags = make(map[string]int)


	for i, s := range input {
		bags[name(s)] = i
	}

	dump()

	return 0
}


func name(rule string) string {
	split := strings.Split(rule, " ")
	return split[0] + " " + split[1]
}

func carry(rule string) map[string]int {
	out := make(map[string]int)
	contain := strings.Split(rule, "contain ")[1]
	for _, bag := range strings.Split(contain, ", ") {
		if bag == "no other bags." {
			continue
		}
		s := strings.Split(bag, " ")
		bagName := s[1] + " " + s[2]
		out[bagName] = Int(s[0])
	}

	return out
}

func dump() {
	for name, id := range bags {
		Printfln("%s %d", name, id)
	}
}

func Solve2(input []string) int {
	return 0
}
