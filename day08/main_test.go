package main

import (
	. "aoc"
	"testing"
)

func Test1(test *testing.T) {
	s :=
`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	AssertEq(test, Solve1(Lines(s)), 5)
}

func Test2(test *testing.T) {

}
