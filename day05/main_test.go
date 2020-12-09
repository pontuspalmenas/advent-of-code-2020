package main

import (
	. "aoc"
	"fmt"
	"testing"
)

func TestSlask(t *testing.T) {
	x := make([]int, 128)
	for i, _ := range x {
		x[i] = i
	}
	fmt.Println(foo("FBFBBFFRLR", x))
}

func foo(s string, n []int) int {
	if len(s) == 1 {
		return n[0]
	}
	if s[0] == 'F' {
		return foo(s[1:], n[:len(n)/2])
	}
	return foo(s[1:], n[len(n)/2:])

}

func Test1(test *testing.T) {
	tt := []TestTable{
		{In: []string{"FBFBBFFRLR"}, Out: 357},
		{In: []string{"BFFFBBFRRR"}, Out: 567},
		{In: []string{"FFFBBBFRRR"}, Out: 119},
		{In: []string{"BBFFBBFRLL"}, Out: 820},
	}

	for _, t := range tt {
		AssertEq(test, Solve1(t.In.([]string)), t.Out)
	}
}

func Test2(test *testing.T) {
	tt := []TestTable{
		{In: []string{}, Out: 0},
	}

	for _, t := range tt {
		AssertEq(test, Solve2(t.In.([]string)), t.Out)
	}
}
