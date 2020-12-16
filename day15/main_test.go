package main

import (
	. "aoc"
	"testing"
)

func Test1(t *testing.T) {
	AssertEq(t, Solve1("0,3,6"), 436)
}

func Test2(test *testing.T) {
	AssertEq(test, Solve2(""), 0)
}
