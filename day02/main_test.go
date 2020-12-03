package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	s := `
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`
	AssertEq(test, Solve1(s), 2)
}

func Test2(test *testing.T) {
	s := `
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

	AssertEq(test, Solve2(s), 1)
}
