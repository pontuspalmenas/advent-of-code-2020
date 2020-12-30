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
`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

	var expect int64 = 208
	AssertEq(test, Solve2(Lines(s)), expect)
}

func TestAddresses(t *testing.T) {
	addr := "000000000000000000000000000000101010"
	mask := "000000000000000000000000000000X1001X"
	floating := addresses(addr, mask)
	AssertEq(t, len(floating), 4)
	AssertEq(t, contains(floating, "000000000000000000000000000000011010"), true)
	AssertEq(t, contains(floating, "000000000000000000000000000000011011"), true)
	AssertEq(t, contains(floating, "000000000000000000000000000000111010"), true)
	AssertEq(t, contains(floating, "000000000000000000000000000000111011"), true)
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
