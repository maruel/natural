// Copyright 2018 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package natural_test

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/maruel/natural"
)

func Example() {
	items := []string{
		"gpio10",
		"gpio1",
		"gpio20",
	}
	slices.SortFunc(items, natural.Compare)
	fmt.Println(strings.Join(items, "\n"))
	// Output:
	// gpio1
	// gpio10
	// gpio20
}

func ExampleLess() {
	// The old way to sort before Go 1.23. It is recommended to use the new slices standard package with Compare.
	items := []string{
		"gpio10",
		"gpio1",
		"gpio20",
	}
	sort.Sort(natural.StringSlice(items))
	fmt.Println(strings.Join(items, "\n"))
	// Output:
	// gpio1
	// gpio10
	// gpio20
}

func ExampleCompare() {
	items := []string{
		"gpio10",
		"gpio1",
		"gpio20",
	}
	slices.SortFunc(items, natural.Compare)
	fmt.Println(strings.Join(items, "\n"))
	// Output:
	// gpio1
	// gpio10
	// gpio20
}
