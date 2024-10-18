package benchmarkfun_test

import (
	"reflect"
	"testing"

	"github.com/willemschots/benchmarkfun"
)

func Test_StructMapSet_Intersection(t *testing.T) {
	for name, tc := range tests() {
		t.Run(name, func(t *testing.T) {
			s1 := benchmarkfun.NewStructMapSet(tc.s1...)
			s2 := benchmarkfun.NewStructMapSet(tc.s2...)
			want := benchmarkfun.NewStructMapSet(tc.want...)

			got := s1.Intersection(s2)

			if !reflect.DeepEqual(got, want) {
				t.Fatalf("want\n%+v\ngot\n%+v\n", want, got)
			}
		})
	}
}

var structMapSetGlobal int

func Benchmark_StructMapSet(b *testing.B) {
	for _, tc := range benchmarks() {
		b.Run(tc.name, func(b *testing.B) {
			// set up the sets
			s1 := benchmarkfun.NewStructMapSet(tc.s1...)
			s2 := benchmarkfun.NewStructMapSet(tc.s2...)

			// set up is not part of the benchmark.
			b.ResetTimer()

			// we now need to be careful the compiler
			// doesn't optimize away our benchmark.
			//
			// To do this, we need to assign our results
			// to a global variable. This is slow however,
			// so we first assign to a local variable.

			// local variable
			var result benchmarkfun.StructMapSet

			// the benchmark
			for range b.N {
				result = s1.Intersection(s2)
			}

			// global variable
			structMapSetGlobal = len(result)
		})
	}
}
