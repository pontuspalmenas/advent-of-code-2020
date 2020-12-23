package aoc

import (
	"testing"
)

func TestMin(test *testing.T) {
	tt := []TestTable{
		{In: []int{1, 2}, Out: 1},
		{In: []int{-1, 0}, Out: -1},
		{In: []int{0, -1}, Out: -1},
		{In: []int{1, 1}, Out: 1},
	}

	for _, t := range tt {
		a := t.In.([]int)[0]
		b := t.In.([]int)[1]
		AssertEq(test, Min(a, b), t.Out)
	}
}

func TestMax(test *testing.T) {
	tt := []TestTable{
		{In: []int{1, 2}, Out: 2},
		{In: []int{-1, 0}, Out: 0},
		{In: []int{0, -1}, Out: 0},
		{In: []int{1, 1}, Out: 1},
	}

	for _, t := range tt {
		a := t.In.([]int)[0]
		b := t.In.([]int)[1]
		AssertEq(test, Max(a, b), t.Out)
	}
}

func TestInputInts(test *testing.T) {
	// Checks both comma-separated and line separated formats
	actual := Ints(Input("aoc_test_input.txt"))
	expected := []int{0,1,-1,2,3}
	AssertEq(test, actual, expected)
}

func TestLines(test *testing.T) {
	actual := Lines(Input("aoc_test_input.txt"))
	expected := []string{"0,1,-1","2","3"}
	AssertEq(test, actual, expected)
}

func TestSscanf(test *testing.T) {
	var n1 int
	var n2 int
	var s1 rune
	var s2 string
	Sscanf("7-11 m: mmmmmmsmmmmm", "%d-%d %c: %s", &n1, &n2, &s1, &s2)

	AssertEq(test, n1, 7)
	AssertEq(test, n2, 11)
	AssertEq(test, s1, 'm')
	AssertEq(test, s2, "mmmmmmsmmmmm")
}

func TestPrintfln(t *testing.T) {
	Printfln("%s %d", "hej", 555)
}

func TestRegex(t *testing.T) {
	match := Regex("^(.*) bags contain (.*)", "vibrant bronze bags contain 4 posh orange bags.")
	AssertEq(t, match[0], "vibrant bronze")
	AssertEq(t, match[1], "4 posh orange bags.")
}

func TestRegexAll(t *testing.T) {
	s :=
`vibrant bronze bags contain 4 posh orange bags.
faded blue bags contain no other bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.`

	match := RegexAll("(.*) bags contain (.*)", s)
	AssertEq(t, match[0][0], "vibrant bronze")
	AssertEq(t, match[0][1], "4 posh orange bags.")

	AssertEq(t, match[1][0], "faded blue")
	AssertEq(t, match[1][1], "no other bags.")

	AssertEq(t, match[2][0], "dark olive")
	AssertEq(t, match[2][1], "3 faded blue bags, 4 dotted black bags.")
}

func TestManhattan(t *testing.T) {
	p1 := Point{X: 3, Y: 4}
	p2 := Point{X: 5, Y: 6}
	actual := Manhattan(p1, p2)
	AssertEq(t, actual, 4)
}

func TestSplitByComma(t *testing.T) {
	s := "foo, bar,baz" // handle both with and without whitespace
	AssertEq(t, SplitByComma(s), []string{"foo","bar","baz"})
}
