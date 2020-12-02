package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	tt := []TestTable{
		{In: []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, Out: 2},
	}

	for _, t := range tt {
		AssertEq(test, Solve1(t.In.([]string)), t.Out)
	}
}

func Test2(test *testing.T) {
	tt := []TestTable{
		{In: []string{}, Out: 2},
	}

	for _, t := range tt {
		AssertEq(test, Solve2(t.In.([]string)), t.Out)
	}
}
