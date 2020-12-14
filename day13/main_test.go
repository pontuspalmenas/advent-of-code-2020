package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	s:=
`939
7,13,x,x,59,x,31,19`

	AssertEq(test, Solve1(Lines(s)), 295)
}

func Test2(test *testing.T) {
	s:=
``

	AssertEq(test, Solve2(Lines(s)), 0)
}
