package types

import (
	"aoc"
	"testing"
)

func TestSet_BasicOperations(t *testing.T) {
	set := NewStringSet()
	set.Add("a")
	set.Add("b")
	aoc.AssertEq(t, set.Size(), 2)
	aoc.AssertEq(t, set.Contains("a"), true)
	aoc.AssertEq(t, set.Contains("b"), true)

	set.Remove("a")
	aoc.AssertEq(t, set.Size(), 1)
	aoc.AssertEq(t, set.Contains("a"), false)
	aoc.AssertEq(t, set.Contains("b"), true)

	slice := []string{"a", "b", "c", "d"}
	aoc.AssertEq(t, FromStringSlice(slice).ToSlice(), slice)
}

func TestSet_Left(t *testing.T) {
	a := FromStringSlice([]string{"a", "b", "c"})
	b := FromStringSlice([]string{"c", "d", "e"})
	left := a.Left(b)

	aoc.AssertEq(t, left.Size(), 2)
	aoc.AssertEq(t, left.Contains("a"), true)
	aoc.AssertEq(t, left.Contains("b"), true)
}

func TestSet_Union(t *testing.T) {
	a := FromStringSlice([]string{"a", "b", "c"})
	b := FromStringSlice([]string{"c", "d", "e"})
	union := a.Union(b)

	aoc.AssertEq(t, union.Size(), 5)
	aoc.AssertEq(t, union.Contains("a"), true)
	aoc.AssertEq(t, union.Contains("b"), true)
	aoc.AssertEq(t, union.Contains("c"), true)
	aoc.AssertEq(t, union.Contains("d"), true)
	aoc.AssertEq(t, union.Contains("e"), true)
}
