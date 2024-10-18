# Go/Golang benchmarking example

This package contains examples of Go benchmarks. The code in this package is part of an [article on benchmarks in Go](https://www.willem.dev/articles/benchmarks-performance-testing/).

This package benchmarks `Intersection` methods on three set implementations:
- `map[string]bool`: Sets as maps of booleans.
- `map[string]struct{}`: Sets as maps of empty structs.
- `[]string`: Sets as slices.

`string` is used as the element type for all sets.

The goal is to learn how to write benchmarks, not to create the most complete or versatile set implementation.

If you're looking for more information on Go sets, check out [this article on sets in Go](https://www.willem.dev/articles/sets-in-golang/).

## Unit tests

To make sure all sets function correctly, unit tests are included as well. These can be found in `benchmarkfun_test.go`.
