package benchmarkfun

// SliceSet implements a set using a slice.
type SliceSet []string

// NewSliceSet creates a new set, if the same value
// is provided is multiple times NewBoolMapSet will panic. This
// is considered a programmer error.
func NewSliceSet(vals ...string) SliceSet {
	set := make(SliceSet, 0, len(vals))
	for _, newV := range vals {
		for _, v := range set {
			if newV == v {
				panic("set cannot contain the same value twice value:" + v)
			}
		}
		set = append(set, newV)
	}

	return set
}

// Intersection finds the intersection of s1 and s2.
func (s1 SliceSet) Intersection(s2 SliceSet) SliceSet {
	result := make(SliceSet, 0)
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				result = append(result, v1)
			}
		}
	}

	return result
}
