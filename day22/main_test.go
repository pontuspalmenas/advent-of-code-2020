package main

import (
	. "aoc"
	"testing"
)

func Test1(t *testing.T) {
	s:=
`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

	AssertEq(t, Solve1(s), 306)
}

func Test2(t *testing.T) {
	s:=
`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

	AssertEq(t, Solve2(s), 291)
}
