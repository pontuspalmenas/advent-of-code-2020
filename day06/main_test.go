package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	in :=
`abc

a
b
c

ab
ac

a
a
a
a

b`

	AssertEq(test, Solve1(in), 11)
}

func Test2(test *testing.T) {
	in :=
`abc

a
b
c

ab
ac

a
a
a
a

b`

	AssertEq(test, Solve2(in), 6)
}
