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
