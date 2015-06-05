// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set_test

import (
	"fmt"
	"sort"

	"github.com/xtgo/set"
)

func ExampleApply() {
	sets := []sort.IntSlice{
		{1, 3, 5, 7, 9},  // odds from 1
		{3, 5, 7, 9, 11}, // odds from 3
		{5, 10, 15, 20},  // 5-multiples
		{2, 3, 5, 7, 11}, // primes
	}

	pivots := make([]int, len(sets))
	var orig, data sort.IntSlice

	// concatenate the sets together for use with the set package
	for i, set := range sets {
		pivots[i] = len(set)
		orig = append(orig, set...)
	}

	// transform set sizes into pivots
	pivots = set.Pivots(pivots...)

	tasks := []struct {
		name string
		op   set.Op
	}{
		{"union", set.Union},
		{"inter", set.Inter},
		{"sdiff", set.SymDiff},
	}

	for _, task := range tasks {
		// make a copy of the original data (Apply rearranges the input)
		data = append(data[:0], orig...)
		size := set.Apply(task.op, data, pivots)
		data = data[:size]
		fmt.Printf("%s: %v\n", task.name, data)
	}

	// Output:
	// union: [1 2 3 5 7 9 10 11 15 20]
	// inter: [5]
	// sdiff: [1 2 3 7 10 15 20]
}

func ExampleApply_diff() {
	// a -  b - c - d  cannot be used with Apply (Diff is non-associative)
	// a - (b + c + d) equivalent, using Apply (Union is associative)

	sets := []sort.IntSlice{
		{0, 2, 4, 6, 8, 10},  // positive evens
		{0, 1, 2, 3, 5, 8},   // set of fibonacci numbers
		{5, 10, 15},          // positive 5-multiples
		{2, 3, 5, 7, 11, 13}, // primes
	}

	var data sort.IntSlice

	// for use with (b + c + d)
	exprsets := sets[1:]
	pivots := make([]int, len(exprsets))

	// concatenate the sets together for use with the set package
	for i, set := range exprsets {
		pivots[i] = len(set)
		data = append(data, set...)
	}

	// transform set sizes into pivots
	pivots = set.Pivots(pivots...)

	// calculate x = (b + c + d)
	size := set.Apply(set.Union, data, pivots)

	// concatenate the result to the end of a
	data = append(sets[0], data[:size]...)

	// calculate a - x
	size = set.Diff(data, len(sets[0]))
	data = data[:size]

	fmt.Println("diff:", data)

	// Output:
	// diff: [4 6]
}
