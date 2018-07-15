// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/xtgo/set"
	"github.com/xtgo/set/internal"
	ss "github.com/xtgo/set/internal/sliceset"
	td "github.com/xtgo/set/internal/testdata"
)

func Example() {
	s := set.Strings([]string{"alpha", "gamma", "alpha"})
	fmt.Println("set:", s)

	s = set.StringsDo(set.Union, s, "beta")
	fmt.Println("set + beta:", s)

	fmt.Println(s, "contains any [alpha delta]:",
		set.StringsChk(set.IsInter, s, "alpha", "delta"))

	fmt.Println(s, "contains all [alpha delta]:",
		set.StringsChk(set.IsSuper, s, "alpha", "delta"))

	// Output:
	// set: [alpha gamma]
	// set + beta: [alpha beta gamma]
	// [alpha beta gamma] contains any [alpha delta]: true
	// [alpha beta gamma] contains all [alpha delta]: false
}

func ExampleUniq() {
	data := sort.IntSlice{5, 7, 3, 3, 5}

	sort.Sort(data)     // sort the data first
	n := set.Uniq(data) // Uniq returns the size of the set
	data = data[:n]     // trim the duplicate elements

	fmt.Println(data)
	// Output: [3 5 7]
}

func ExampleInter() {
	data := sort.IntSlice{3, 5, 7} // create a set (it must be sorted)
	pivot := len(data)             // store the length of our first set

	data = append(data, 1, 3, 5)   // append a second set (which also must be sorted)
	size := set.Inter(data, pivot) // perform set intersection

	// trim data to contain just the result set
	data = data[:size]

	fmt.Println("inter:", data)

	// Output:
	// inter: [3 5]
}

func ExampleIsSuper() {
	data := sort.StringSlice{"b", "c", "d"} // create a set (it must be sorted)
	pivot := len(data)                      // store the length of our first set

	data = append(data, "c", "d")         // append a second set (which also must be sorted)
	contained := set.IsSuper(data, pivot) // check the first set is a superset of the second

	fmt.Printf("%v superset of %v = %t\n", data[:pivot], data[pivot:], contained)

	data = data[:pivot] // trim off the second set

	data = append(data, "s")             // append a new singleton set to compare against
	contained = set.IsSuper(data, pivot) // check for membership

	fmt.Printf("%v superset of %v = %t\n", data[:pivot], data[pivot:], contained)

	// Output:
	// [b c d] superset of [c d] = true
	// [b c d] superset of [s] = false
}

func ExampleIsInter() {
	data := sort.StringSlice{"b", "c", "d"} // create a set (it must be sorted)
	pivot := len(data)                      // store the length of our first set

	data = append(data, "d", "z")         // append a second set (which also must be sorted)
	contained := set.IsInter(data, pivot) // check the first set is a superset of the second

	fmt.Printf("%v intersects %v = %t\n", data[:pivot], data[pivot:], contained)

	data = data[:pivot] // trim off the second set

	data = append(data, "s")             // append a new singleton set to compare against
	contained = set.IsInter(data, pivot) // check for membership

	fmt.Printf("%v intersects %v = %t\n", data[:pivot], data[pivot:], contained)

	// Output:
	// [b c d] intersects [d z] = true
	// [b c d] intersects [s] = false
}

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

func BenchmarkApply_256x64K(b *testing.B) {
	benchApply(b, set.Apply, td.Rand(256, td.Large))
}

func BenchmarkApplyInSeq_256x64K(b *testing.B) {
	benchApply(b, ApplyInSeq, td.Rand(256, td.Large))
}

func datalen(sets [][]int) int {
	sum := 0
	for _, set := range sets {
		sum += len(set)
	}
	return sum
}

func mkpivots(sets [][]int) []int {
	lengths := make([]int, len(sets))
	for i, set := range sets {
		lengths[i] = len(set)
	}
	return set.Pivots(lengths...)
}

func mkdata(sets [][]int) ss.Set {
	data := make(ss.Set, 0, datalen(sets))
	for _, set := range sets {
		data = append(data, set...)
	}
	return data
}

type applyFunc func(set.Op, sort.Interface, []int) int

func benchApply(b *testing.B, fn applyFunc, sets [][]int) {
	pivots := mkpivots(sets)
	src := mkdata(sets)
	data := make(ss.Set, len(src))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(data, src)
		fn(set.Inter, data, pivots)
	}
}

func ApplyInSeq(op set.Op, data sort.Interface, pivots []int) (size int) {
	n := len(pivots)
	if n <= 1 {
		return data.Len()
	}

	k := pivots[n-1]
	// we're doing one less iteration to account for a single lookbehind
	for ii := range pivots[1:] {
		// reverse iteration, for locality (results propagate toward zero)
		ii = len(pivots) - ii - 2
		i := 0 // start
		if ii > 0 {
			i = pivots[ii-1]
		}
		j := pivots[ii] // pivot
		b := internal.BoundSpan{data, internal.Span{i, k}}
		n := op(b, j-i)
		k = i + n
	}
	return k
}

func TestApplyInSeq(t *testing.T) {
	sets := td.Rand(8, td.Small)
	pivots := mkpivots(sets)
	data1 := mkdata(sets)
	data2 := mkdata(sets)
	size1 := ApplyInSeq(set.Inter, data2, pivots)
	size2 := set.Apply(set.Inter, data1, pivots)
	data1 = data1[:size1]
	data2 = data2[:size2]

	if size1 != size2 {
		t.Errorf("ApplyInSeq call got %d, want %d", size1, size2)
	}
	if !data1.IsEqual(data2) {
		t.Fatalf("ApplyInSeq output differed:\n   got: %v\n  want: %v", data1, data2)
	}
}
