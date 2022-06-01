// Package benchmark provides ...
package benchmark

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	facette "github.com/facette/natsort"
	"github.com/maruel/natural"
)

func BenchmarkNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if less("a01a2", "a01a01") {
			b.Fatal("unexpected result")
		}
	}
}

//go:noinline
func less(a, b string) bool {
	return a < b
}

var dre = regexp.MustCompile(`\d+`)

func pad(str string) string {
	f := func(match []byte) []byte {
		n, _ := strconv.ParseInt(string(match), 10, 64)
		return []byte(fmt.Sprintf("%010d", n))
	}

	return string(dre.ReplaceAllFunc([]byte(str), f))
}

//go:noinline
func lessDummy(a, b string) bool {
	return pad(a) < pad(b)
}

func BenchmarkLessStringOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if natural.Less("abcd", "abc") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessStringOnlyFacette(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if facette.Compare("abcd", "abc") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessStringOnlyDummy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if lessDummy("abcd", "abc") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if natural.Less("10", "2") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessDigitsFacette(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if facette.Compare("10", "2") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessDigitsDummy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if lessDummy("10", "2") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessStringDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if natural.Less("a10", "a2") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessStringDigitsFacette(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if facette.Compare("a10", "a2") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessStringDigitsDummy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if lessDummy("a10", "a2") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessDigitsTwoGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if natural.Less("a01a2", "a01a01") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessDigitsTwoGroupsFacette(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if facette.Compare("a01a2", "a01a01") {
			b.Fatal("unexpected result")
		}
	}
}

func BenchmarkLessDigitsTwoGroupsDummy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if lessDummy("a01a2", "a01a01") {
			b.Fatal("unexpected result")
		}
	}
}
