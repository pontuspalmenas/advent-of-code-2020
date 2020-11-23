package template

import (
	"aoc"
	"testing"
)

func Test1(test *testing.T) {
	tt := []aoc.TestTable{
		{In: 0, Out: 0},
	}

	for _, t := range tt {
		in := []int{t.In.(int)}
		aoc.AssertEq(test, Solve1(in), t.Out)
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
