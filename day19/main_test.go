package main

import (
	. "aoc"
	"regexp"
	"testing"
)

func Test1(t *testing.T) {
	s:=
`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

	AssertEq(t, Solve1(SplitByEmptyNewline(s)), 2)
}

func TestRegx(t *testing.T) {
	in:=
`ababbb
bababa
abbbab
aaabbb
aaaabbb`

	r := regexp.MustCompile(`^(?:(?:a)(?:(?:(?:a)(?:a)|(?:b)(?:b))(?:(?:a)(?:b)|(?:b)(?:a))|(?:(?:a)(?:b)|(?:b)(?:a))(?:(?:a)(?:a)|(?:b)(?:b)))(?:b))$`)

	for _, s := range Lines(in) {
		if r.MatchString(s) {
			Printfln("valid: %s", s)
		} else {
			Printfln("invalid: %s", s)
		}
	}

}

func Test2(t *testing.T) {
	s:=
``

	AssertEq(t, Solve2(SplitByEmptyNewline(s)), 0)
}
