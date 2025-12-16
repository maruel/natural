// Copyright 2018 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package natural

import (
	"fmt"
	"testing"
)

func TestCompareLess(t *testing.T) {
	data := [][2]string{
		{"", "a"},
		{"a", "b"},
		{"a", "aa"},
		{"a0", "a1"},
		{"a0", "a00"},
		{"a00", "a01"},
		{"a01", "a1"},
		{"a01", "a2"},
		{"a01x", "a2x"},
		// Only the last number matters.
		{"a0b00", "a00b1"},
		{"a0b00", "a00b01"},
		{"a00b0", "a0b00"},
		{"a00b00", "a0b01"},
		{"a00b00", "a0b1"},
		// Number larger than uint64 max - this will fail due to overflow
		// uint64 max is 18446744073709551615, these numbers exceed that
		{"a99999999999999999999", "a100000000000000000000"},
	}
	for _, l := range data {
		t.Run(fmt.Sprintf("%s-%s", l[0], l[1]), func(t *testing.T) {
			if res := Compare(l[0], l[1]); res >= 0 {
				t.Fatalf("Compare(%q, %q) returned >=0: %d", l[0], l[1], res)
			}
		})
	}
}

func TestCompareGreater(t *testing.T) {
	data := [][2]string{
		{"a", ""},
		{"aa", "a"},
		{"b", "a"},
		{"a01", "a00"},
		{"a2", "a01"},
		{"a2x", "a01x"},
		{"a00b00", "a0b0"},
		{"a00b01", "a0b00"},
		{"10", "2"},
	}
	for _, l := range data {
		t.Run(fmt.Sprintf("%s-%s", l[0], l[1]), func(t *testing.T) {
			if res := Compare(l[0], l[1]); res <= 0 {
				t.Fatalf("Compare(%q, %q) returned <=0: %d", l[0], l[1], res)
			}
		})
	}
}

func TestCompareEqual(t *testing.T) {
	data := [][2]string{
		{"a", "a"},
		{"a01", "a01"},
		{"a1", "a1"},
		{"a00b00", "a0b00"},
		{"a0b00", "a0b00"},
		// Large numbers with leading zeros are equal when there's trailing data on both sides
		{"a00000000000000000000001x", "a1x"},
		{"a099999999999999999999x", "a99999999999999999999x"},
	}
	for _, l := range data {
		t.Run(fmt.Sprintf("%s-%s", l[0], l[1]), func(t *testing.T) {
			if res := Compare(l[0], l[1]); res != 0 {
				t.Fatalf("Compare(%q, %q) returned !=0: %d", l[0], l[1], res)
			}
		})
	}
}

func TestCompareLargeNumbers(t *testing.T) {
	// Additional tests specifically for numbers larger than uint64
	data := [][3]interface{}{
		// [a, b, expected_sign] where expected_sign: -1 (a<b), 0 (a==b), 1 (a>b)
		{"a99999999999999999999", "a100000000000000000000", -1},
		{"a123456789012345678901234567890", "a123456789012345678901234567891", -1},
		{"a999999999999999999999", "a1000000000000000000000", -1},
		{"a20000000000000000000", "a100000000000000000000", -1},
		{"a100000000000000000000", "a20000000000000000000", 1},
		{"a1000000000000000000000", "a999999999999999999999", 1},
		{"a100000000000000000000", "a100000000000000000000", 0},
		// Leading zeros with trailing data on both sides
		{"a099999999999999999999x", "a99999999999999999999x", 0},
		{"a00000000000000000000001x", "a1x", 0},
	}
	for _, l := range data {
		a := l[0].(string)
		b := l[1].(string)
		expected := l[2].(int)
		t.Run(fmt.Sprintf("%s-%s", a, b), func(t *testing.T) {
			res := Compare(a, b)
			if (expected < 0 && res >= 0) || (expected > 0 && res <= 0) || (expected == 0 && res != 0) {
				t.Fatalf("Compare(%q, %q) returned %d, expected sign: %d", a, b, res, expected)
			}
		})
	}
}

// BenchmarkLessDigitsTwoGroupsNative benchmarks calling a function that just
// does a < b without taking into account the digits, using the same string as
// BenchmarkLessDigitsTwoGroups.
func BenchmarkLessDigitsTwoGroupsNative(b *testing.B) {
	b.ReportAllocs()
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

// BenchmarkLessDigitsTwoGroups compares "a01a2" and "a01a01".
func BenchmarkLessDigitsTwoGroups(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if Less("a01a2", "a01a01") {
			b.Fatal("unexpected result")
		}
	}
}

// BenchmarkLessStringOnly doesn't contain digits.
func BenchmarkLessStringOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if Less("abcd", "abc") {
			b.Fatal("unexpected result")
		}
	}
}

// BenchmarkLessDigitsOnly contains only digits.
func BenchmarkLessDigitsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if Less("10", "2") {
			b.Fatal("unexpected result")
		}
	}
}

// BenchmarkLess10Blocks benchmark a mix of 10 strings and digits.
func BenchmarkLess10Blocks(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if Less("a01a01a01a01a01a01a01a01a01a2", "a01a01a01a01a01a01a01a01a01a01") {
			b.Fatal("unexpected result")
		}
	}
}
