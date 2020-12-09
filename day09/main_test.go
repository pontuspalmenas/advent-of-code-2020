package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
s :=
`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

	AssertEq(test, Solve1(Ints(s), 5), 127)

}

func Test2(test *testing.T) {
	tt := []TestTable{
		{In: 0, Out: 0},
	}

	for _, t := range tt {
		in := []int{t.In.(int)}
		AssertEq(test, Solve2(in), t.Out)
	}
}
