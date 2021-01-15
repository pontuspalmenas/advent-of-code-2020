package main

import (
	. "aoc"
	"testing"
)

func Test1(t *testing.T) {
	AssertEq(t, Solve1("389125467"), "67384529")
}

func Test2(t *testing.T) {
	AssertEq(t, Solve2("389125467"), 149245887792)
}
