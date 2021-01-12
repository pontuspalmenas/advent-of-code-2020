package main

import (
	. "aoc"
	"testing"
)

func Test1(t *testing.T) {
	s:=
`389125467`

	AssertEq(t, Solve1(Lines(s)), 92658374)
}

func Test2(t *testing.T) {
	s:=
``

	AssertEq(t, Solve2(Lines(s)), 0)
}
