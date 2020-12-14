package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	s:=
`19
16
15
12
11
10
7
6
5
4
1`
	input := Ints(s)
	AssertEq(test, Solve1(input), 35)
}

func Test2(test *testing.T) {
	s:=
``

	AssertEq(test, Solve2(Ints(s)), 0)
}
