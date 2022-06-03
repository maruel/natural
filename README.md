# natural

Yet another natural sort, with 100% test coverage and a benchmark. It **does not
allocate memory**, doesn't depend on package `sort` and hence doesn't depend on
`reflect`. It is optimized for speed.

[![Go
Reference](https://pkg.go.dev/badge/github.com/maruel/natural.svg)](https://pkg.go.dev/github.com/maruel/natural)
[![codecov](https://codecov.io/gh/maruel/natural/branch/main/graph/badge.svg?token=iQg8Y62BBg)](https://codecov.io/gh/maruel/natural)


## Benchmarks

On Go 1.18.3.

```
$ go test -bench=. -cpu 1
goos: linux
goarch: amd64
pkg: github.com/maruel/natural
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkLessDigitsTwoGroupsNative 329085624     3.628 ns/op   0 B/op   0 allocs/op
BenchmarkLessDigitsTwoGroups        31792579    38.07 ns/op    0 B/op   0 allocs/op
BenchmarkLessStringOnly            147547137     8.127 ns/op   0 B/op   0 allocs/op
BenchmarkLessDigitsOnly             65866989    17.92 ns/op    0 B/op   0 allocs/op
BenchmarkLess10Blocks                5997386   198.3 ns/op     0 B/op   0 allocs/op
```

On a Raspberry Pi 3:

```
$ go test -bench=. -cpu 1
goos: linux
goarch: arm
pkg: github.com/maruel/natural
BenchmarkLessDigitsTwoGroupsNative  13044535    85.86 ns/op    0 B/op   0 allocs/op
BenchmarkLessDigitsTwoGroups         1576779   751.7 ns/op     0 B/op   0 allocs/op
BenchmarkLessStringOnly              8470698   141.5 ns/op     0 B/op   0 allocs/op
BenchmarkLessDigitsOnly              3674454   326.4 ns/op     0 B/op   0 allocs/op
BenchmarkLess10Blocks                 314845  3821 ns/op       0 B/op   0 allocs/op
```

Coverage:

```
$ go test -cover
PASS
coverage: 100.0% of statements
ok     github.com/maruel/natural       0.012s
```

