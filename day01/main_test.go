package main

import (
	"aoc"
	"testing"
)

func Test1(test *testing.T) {
	tt := []aoc.TestTable{
		{In: []int{1721,
			979,
			366,
			299,
			675,
			1456}, Out: 514579},
	}

	for _, t := range tt {
		aoc.AssertEq(test, Solve1(t.In.([]int)), t.Out)
	}
}

func Test2(test *testing.T) {
	tt := []aoc.TestTable{
		{In: 0, Out: 0},
	}

	for _, t := range tt {
		in := []int{t.In.(int)}
		aoc.AssertEq(test, Solve2(in), t.Out)
	}
}
