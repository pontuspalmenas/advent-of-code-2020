package main

import (
	. "aoc"
	"testing"
)

func Test1(t *testing.T) {
	s:=
`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

	AssertEq(t, Solve1(Lines(s)), 5)
}

func Test2(t *testing.T) {
	s:=
``

	AssertEq(t, Solve2(Lines(s)), 0)
}
