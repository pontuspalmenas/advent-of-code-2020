package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	s:=
`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

	var expect int64 = 165
	AssertEq(test, Solve1(Lines(s)), expect)
}

func Test2(test *testing.T) {
	s:=
``

	AssertEq(test, Solve2(Lines(s)), 0)
}
