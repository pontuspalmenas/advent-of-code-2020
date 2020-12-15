package main

import (
	. "aoc"
	"testing"
)

const small =
`19
16
15
12
11
10
7
6
5
4
1`
const large =
`49
48
47
46
45
42
39
38
35
34
33
32
31
28
25
24
23
20
19
18
17
14
11
10
9
8
7
4
3
2
1`

func Test1Small(test *testing.T) {
	input := Ints(small)
	AssertEq(test, Solve1(input), 35)
}

func Test1Large(test *testing.T) {
	input := Ints(large)
	AssertEq(test, Solve1(input), 220)
}

func Test2Small(test *testing.T) {
	AssertEq(test, Solve2(Ints(small)), 0)
}

func Test2Large(test *testing.T) {
	AssertEq(test, Solve2(Ints(large)), 0)
}
