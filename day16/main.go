package main

import (
	. "aoc"
	. "aoc/types"
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

	tickets, _ := sortTickets(nearby, fields)

	println(len(tickets))
	//mapped := make(map[string]int)
	var unmapped []string

	// find candidates
	candidates := make(map[int]*StringSet)
	for i := 0; i <= len(tickets[0])-1; i++ {
		candidates[i] = NewStringSet()
	}

	for _, t := range tickets {
		for i, v := range t {
			candidates[i] = candidates[i].Union(candidatesForIndex(fields, v))
		}
	}
	// debug
	for k, v := range candidates {
		Printfln("%d -> %v", k, v)
	}

	for _, f := range fields {
		unmapped = append(unmapped, f.name)
	}

	return 0
}

func candidatesForIndex(fields []Field, value int) *StringSet {
	set := NewStringSet()
	for _, f := range fields {
		if inRange(f, value) {
			set.Add(f.name)
		}
	}
	return set
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
