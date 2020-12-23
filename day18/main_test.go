package main

import (
	. "aoc"
	"testing"
)

func Test1(t *testing.T) {
	s:=
`2 * 3 + (4 * 5)`
	AssertEq(t, Solve1(Lines(s)), 26)
}

func Test2(t *testing.T) {
	s:=
``

	AssertEq(t, Solve2(Lines(s)), 0)
}
