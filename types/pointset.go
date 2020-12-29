package types

import . "aoc"

type PointSet struct {
	m map[Point]struct{}
}

func (s *PointSet) Add(v Point) {
	s.m[v] = struct{}{}
}

func (s *PointSet) Remove(v Point) {
	delete(s.m, v)
}

func (s *PointSet) Contains(v Point) bool {
	_, ok := s.m[v]
	return ok
}

func (s *PointSet) Size() int {
	return len(s.m)
}

func NewPointSet() *PointSet {
	s := &PointSet{}
	s.m = make(map[Point]struct{})
	return s
}

func (s *PointSet) ToSlice() (out []Point) {
	for m := range s.m {
		out = append(out, m)
	}
	return out
}

