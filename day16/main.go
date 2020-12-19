package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := Input("day16/input.txt")
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

type Range struct {
	lower int
	upper int
}

type Field struct {
	name string
	range1 Range
	range2 Range
}

type Ticket []int

func Solve1(input string) int {
	split := SplitByEmptyNewline(input)
	rules := Lines(split[0])
	nearby := tickets(Lines(split[2])[1:])
	var fields []Field

	for _, r := range rules {
		fields = append(fields, parseField(r))
	}

	_, invalid := sortTickets(nearby, fields)
	return invalid
}

func Solve2(input string) int {
	split := SplitByEmptyNewline(input)
	rules := Lines(split[0])
	nearby := tickets(Lines(split[2])[1:])
	var fields []Field

	for _, r := range rules {
		fields = append(fields, parseField(r))
	}

	valid, _ := sortTickets(nearby, fields)

	var freeColumns []int
	var mapped map[string]int
	var revMap map[int]string



	return 0
}

func sortTickets(tickets []Ticket, fields []Field) (valid []Ticket, sumInvalid int) {
	for _, t := range tickets {
		validTicket := true
		for _, v := range t {
			validValue := false
			for _, f := range fields {
				if inRange(f, v) {
					validValue = true
					break
				}
			}
			if !validValue {
				sumInvalid += v
				validTicket = false
			}
		}
		if validTicket {
			valid = append(valid, t)
		}
	}
	return valid, sumInvalid
}

func tickets(ss []string) []Ticket {
	var tt []Ticket
	for _, s := range ss {
		tt = append(tt, Ints(s))
	}

	return tt
}

func inRange(f Field, v int) bool {
	return (v >= f.range1.lower && v <= f.range1.upper) ||
		(v >= f.range2.lower && v <= f.range2.upper)
}

func parseField(s string) Field {
	parts := Regex(`^(\w+ ?\w+): (\d+)-(\d+) or (\d+)-(\d+)`, s)
	return Field{
		name: parts[0],
		range1: Range{lower: Int(parts[1]), upper: Int(parts[2])},
		range2: Range{lower: Int(parts[3]), upper: Int(parts[4])},
	}
}
