package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	s:=
`F10
N3
F7
R90
F11`

	AssertEq(test, Solve1(Lines(s)), 25)
}

func Test2(test *testing.T) {
	s:=
`F10
N3
F7
R90
F11`

	AssertEq(test, Solve2(Lines(s)), 286)
}
