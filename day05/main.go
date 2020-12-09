package main

import (
	. "aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := Lines(Input("day05/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []string) int {
	highest := 0

	for _, s := range input {
		r := row(s[:7])
		c := col(s[7:])
		if r*8+c > highest {
			highest = r*8+c
		}
		fmt.Printf("%s: row %d, column %d, seat ID %d\n", s, r, c, r*8+c)
	}

	return highest
}

func row(s string) int {
	s = strings.ReplaceAll(s, "F", "0")
	s = strings.ReplaceAll(s, "B", "1")
	r, _ := strconv.ParseInt(s, 2, 0)
	return int(r)
}

func col(s string) int {
	s = strings.ReplaceAll(s, "L", "0")
	s = strings.ReplaceAll(s, "R", "1")
	r, _ := strconv.ParseInt(s, 2, 0)
	return int(r)
}

func Solve2(input []string) int {
	seats := make(map[int]bool, 994) // magic number
	// get seated
	for _, s := range input {
		r := row(s[:7])
		c := col(s[7:])
		seat := r*8+c
		seats[seat] = true
	}
	// my seat is not in the very back, and not in the very front
	// don't know where, so start going backwards (should probably quit before 994?)
	for i := 994/2; i < 994; i++ {
		if !seats[i] {
			return i
		}
	}

	return 0
}
