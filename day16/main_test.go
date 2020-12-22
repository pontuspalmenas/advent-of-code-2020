package main

import (
	. "aoc"
	"testing"
)

func Test1(t *testing.T) {
	s:=
`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

	AssertEq(t, Solve1(s), 71)
}

func Test2(t *testing.T) {
	s:=
``

	AssertEq(t, Solve2(s), 0)
}
