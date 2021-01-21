package types

import (
	"fmt"
)

type StringSet struct {
	m map[string]struct{}
}

func NewStringSet() *StringSet {
	s := &StringSet{}
	s.m = make(map[string]struct{})
	return s
}

func ToStringSet(s []string) *StringSet {
	out := NewStringSet()
	for _, v := range s {
		out.Add(v)
	}
	return out
}

func (s *StringSet) ToSlice() (out []string) {
	for m := range s.m {
		out = append(out, m)
	}
	return out
}

func (s *StringSet) Add(v string) {
	s.m[v] = struct{}{}
}

func (s *StringSet) Remove(v string) {
	delete(s.m, v)
}

func (s *StringSet) Contains(v string) bool {
	_, ok := s.m[v]
	return ok
}

func (s *StringSet) Size() int {
	return len(s.m)
}

// Left or relative complement
func (s *StringSet) Left(s2 *StringSet) *StringSet {
	out := NewStringSet()
	for v := range s.m {
		if !s2.Contains(v) {
			out.Add(v)
		}
	}
	return out
}

func (s *StringSet) Union(s2 *StringSet) *StringSet {
	out := NewStringSet()
	for v := range s.m {
		out.Add(v)
	}

	for v := range s2.m {
		out.Add(v)
	}
	return out
}

func (s *StringSet) Intersection(s2 *StringSet) *StringSet {
	out := NewStringSet()
	for v := range s.m {
		if s2.Contains(v) {
			out.Add(v)
		}
	}

	return out
}

func (s *StringSet) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
