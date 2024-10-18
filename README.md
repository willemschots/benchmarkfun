# Go/Golang benchmarking example

This package contains examples of Go benchmarks. The code in this package is part of this article on benchmarks in Go.

In the code we compare three different set implementations and benchmark the effort required to get intersections:
- `map[string]bool`: Sets as maps of booleans.
- `map[string]struct{}`: Sets as maps of empty structs.
- `[]string`: Sets as slices.

For all sets we will use `string` as the element type.

The goal is to learn how to write benchmarks, not to create the most complete or versatile set implementation. If you're looking for that, check this article.

## Unit tests

To make sure all sets function correctly, unit tests are included as well. These can be found in `benchmarkfun_test.go`.
