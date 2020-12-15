package main

import (
	. "aoc"
	"testing"
)

const example =
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

func Test1(test *testing.T) {
	AssertEq(test, Solve1(Lines(example)), 37)
}

func Test2(test *testing.T) {
	AssertEq(test, Solve2(Lines(example)), 26)
}
