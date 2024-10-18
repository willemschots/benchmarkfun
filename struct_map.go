package benchmarkfun

// StructMapSet implements a set using a map of empty structs.
type StructMapSet map[string]struct{}

// NewStructMapSet creates a new set, if the same value
// is provided is multiple times NewStructMapSet will panic. This
// is considered a programmer error.
func NewStructMapSet(vals ...string) StructMapSet {
	set := make(StructMapSet, len(vals))
	for _, v := range vals {
		_, ok := set[v]
		if ok {
			panic("set cannot contain the same value twice value:" + v)
		}
		set[v] = struct{}{}
	}
	return set
}

// Intersection finds the intersection of s1 and s2.
func (s1 StructMapSet) Intersection(s2 StructMapSet) StructMapSet {
	result := make(StructMapSet, 0)
	for k := range s1 {
		if _, ok := s2[k]; ok {
			result[k] = struct{}{}
		}
	}
	return result
}
