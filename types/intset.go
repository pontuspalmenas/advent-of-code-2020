package types

import "fmt"

type IntSet struct {
	m map[int]struct{}
}

func NewIntSet() *IntSet {
	s := &IntSet{}
	s.m = make(map[int]struct{})
	return s
}

func FromIntSlice(s []int) *IntSet {
	out := NewIntSet()
	for _, v := range s {
		out.Add(v)
	}
	return out
}

func (s *IntSet) ToSlice() (out []int) {
	for m := range s.m {
		out = append(out, m)
	}
	return out
}

func (s *IntSet) Add(v int) {
	s.m[v] = struct{}{}
}

func (s *IntSet) Remove(v int) {
	delete(s.m, v)
}

func (s *IntSet) Contains(v int) bool {
	_, ok := s.m[v]
	return ok
}

func (s *IntSet) Size() int {
	return len(s.m)
}

// Left or relative complement
func (s *IntSet) Left(s2 *IntSet) *IntSet {
	out := NewIntSet()
	for v := range s.m {
		if !s2.Contains(v) {
			out.Add(v)
		}
	}
	return out
}

func (s *IntSet) Union(s2 *IntSet) *IntSet {
	out := NewIntSet()
	for v := range s.m {
		out.Add(v)
	}

	for v := range s2.m {
		out.Add(v)
	}
	return out
}

func (s *IntSet) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
