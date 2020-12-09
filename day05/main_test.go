package main

import (
	. "aoc"
	"testing"
)

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
