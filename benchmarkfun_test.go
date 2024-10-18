package benchmarkfun_test

import "fmt"

type test struct {
	s1   []string
	s2   []string
	want []string
}

func tests() map[string]test {
	return map[string]test{
		"all empty":               {s1: []string{}, s2: []string{}, want: []string{}},
		"distinct min":            {s1: []string{"a"}, s2: []string{"b"}, want: []string{}},
		"distinct max":            {s1: []string{"a", "b", "c"}, s2: []string{"d", "e", "f"}, want: []string{}},
		"intersection min":        {s1: []string{"a"}, s2: []string{"a"}, want: []string{"a"}},
		"intersection max":        {s1: []string{"a", "b", "c"}, s2: []string{"a", "b", "c"}, want: []string{"a", "b", "c"}},
		"intersection empty str.": {s1: []string{""}, s2: []string{""}, want: []string{""}},
		"intersection s1 > s2":    {s1: []string{"a", "b", "c"}, s2: []string{"b"}, want: []string{"b"}},
		"intersection s1 < s2":    {s1: []string{"b"}, s2: []string{"a", "b", "c"}, want: []string{"b"}},
	}
}

type benchmark struct {
	name string
	s1   []string
	s2   []string
}

func benchmarks() []benchmark {
	vals := func(length int, prefix string) []string {
		v := make([]string, length)
		for i := 0; i < length; i++ {
			v[i] = fmt.Sprintf("%s%d", prefix, i)
		}

		return v
	}

	return []benchmark{
		{
			name: "10, 10 no intersection",
			s1:   vals(10, "a"),
			s2:   vals(10, "b"),
		},
		{
			name: "10, 10 full intersection",
			s1:   vals(10, ""),
			s2:   vals(10, ""),
		},
		{
			name: "100, 100 no intersection",
			s1:   vals(100, "a"),
			s2:   vals(100, "b"),
		},
		{
			name: "100, 100 full intersection",
			s1:   vals(100, ""),
			s2:   vals(100, ""),
		},
		{
			name: "1000, 1000 no intersection",
			s1:   vals(1000, "a"),
			s2:   vals(1000, "b"),
		},
		{
			name: "1000, 1000 full intersection",
			s1:   vals(1000, ""),
			s2:   vals(1000, ""),
		},
	}
}
