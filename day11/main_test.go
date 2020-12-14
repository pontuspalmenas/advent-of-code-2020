package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	s:=
`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	AssertEq(test, Solve1(Lines(s)), 37)
}

func Test2(test *testing.T) {
	s:=
``

	AssertEq(test, Solve2(Lines(s)), 0)
}
