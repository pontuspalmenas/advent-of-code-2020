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

func TestChrAt(test *testing.T) {
	actual := ChrAt("abc123", 2)
	expected := "c"
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

func TestFscanf(test *testing.T) {
	var n1 int
	var n2 int
	var s1 rune
	var s2 string
	Sscanf("aoc_test_input2.txt", "%d-%d %c: %s", &n1, &n2, &s1, &s2)

	AssertEq(test, n1, 7)
	AssertEq(test, n2, 11)
	AssertEq(test, s1, 'm')
	AssertEq(test, s2, "mmmmmmsmmmmm")
}

func TestPrintfln(t *testing.T) {
	Printfln("%s %d", "hej", 555)
}
