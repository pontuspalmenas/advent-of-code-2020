package types

import (
	"aoc"
	"testing"
)

func TestMatrix_Flip(t *testing.T) {
	m := NewMatrix(4,4)
	m.data = [][]int{
		{1,2,3,4},
		{1,2,3,4},
		{1,2,3,4},
		{1,2,3,4}}

	expected := [][]int{
		{4,3,2,1},
		{4,3,2,1},
		{4,3,2,1},
		{4,3,2,1}}

	m.Flip()

	aoc.AssertEq(t, m.data, expected)
}

func TestMatrix_Rotate(t *testing.T) {
	m := NewMatrix(4,4)
	m.data = [][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,0,1,2},
		{3,4,5,6}}

	expected := [][]int{
		{3,9,5,1},
		{4,0,6,2},
		{5,1,7,3},
		{6,2,8,4}}

	m.Rotate()

	aoc.AssertEq(t, m.data, expected)
}
