package main

import (
	. "aoc"
	"testing"
)

func Test1(t *testing.T) {
	s:=
`5764801
17807724`

	AssertEq(t, Solve1(Lines(s)), 14897079)
}
