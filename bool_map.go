package benchmarkfun

// BoolMapSet implements a set using a map of booleans.
type BoolMapSet map[string]bool

// NewBoolMapSet creates a new set, if the same value
// is provided is multiple times NewBoolMapSet will panic. This
// is considered a programmer error.
func NewBoolMapSet(vals ...string) BoolMapSet {
	set := make(BoolMapSet, len(vals))
	for _, v := range vals {
		_, ok := set[v]
		if ok {
			panic("set cannot contain the same value twice value:" + v)
		}
		set[v] = true
	}
	return set
}

// Intersection finds the intersection of s1 and s2.
func (s1 BoolMapSet) Intersection(s2 BoolMapSet) BoolMapSet {
	result := make(BoolMapSet, 0)
	for k := range s1 {
		if _, ok := s2[k]; ok {
			result[k] = true
		}
	}
	return result
}
