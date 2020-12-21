package aoc

import (
	"testing"
)

func TestSet_BasicOperations(t *testing.T) {
	set := NewStringSet()
	set.Add("a")
	set.Add("b")
	AssertEq(t, set.Size(), 2)
	AssertEq(t, set.Contains("a"), true)
	AssertEq(t, set.Contains("b"), true)

	set.Remove("a")
	AssertEq(t, set.Size(), 1)
	AssertEq(t, set.Contains("a"), false)
	AssertEq(t, set.Contains("b"), true)

	slice := []string{"a", "b", "c", "d"}
	AssertEq(t, ToStringSet(slice).ToSlice(), slice)
}

func TestSet_Left(t *testing.T) {
	a := ToStringSet([]string{"a", "b", "c"})
	b := ToStringSet([]string{"c", "d", "e"})
	left := a.Left(b)

	AssertEq(t, left.Size(), 2)
	AssertEq(t, left.Contains("a"), true)
	AssertEq(t, left.Contains("b"), true)
}

func TestSet_Union(t *testing.T) {
	a := ToStringSet([]string{"a", "b", "c"})
	b := ToStringSet([]string{"c", "d", "e"})
	union := a.Union(b)

	AssertEq(t, union.Size(), 5)
	AssertEq(t, union.Contains("a"), true)
	AssertEq(t, union.Contains("b"), true)
	AssertEq(t, union.Contains("c"), true)
	AssertEq(t, union.Contains("d"), true)
	AssertEq(t, union.Contains("e"), true)
}
