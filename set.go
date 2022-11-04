package set

type Set struct {
	m map[string]struct{}
}

func NewSet() *Set {
	return &Set{m: make(map[string]struct{})}
}

func (s *Set) Add(e string) {
	s.m[e] = struct{}{}
}

func (s *Set) Has(e string) bool {
	_, ok := s.m[e]
	return ok
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) Del(e string) {
	delete(s.m, e)
}

func (s *Set) Intersection(s2 *Set) *Set {
	r := NewSet()
	for k := range s.m {
		if s2.Has(k) {
			r.Add(k)
		}
	}
	return r
}

func (s *Set) Dif(s2 *Set) *Set {
	r := NewSet()
	for k := range s.m {
		if !s2.Has(k) {
			r.Add(k)
		}
	}
	return r
}

func (s *Set) ToSlice() []string {
	ret := make([]string, 0, s.Len())
	for k := range s.m {
		ret = append(ret, k)
	}
	return ret
}

func (s *Set) Map() map[string]struct{} {
	return s.m
}
