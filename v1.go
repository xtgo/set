// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package set implements type-safe, non-allocating algorithms that operate
// on ordered sets.
//
// Most functions take a data parameter of type sort.Interface and a pivot
// parameter of type int; data represents two sets covering the ranges
// [0:pivot] and [pivot:Len], each of which is expected to be sorted and
// free of duplicates. sort.Sort may be used for sorting, and Uniq may be
// used to filter away duplicates.
//
// All mutating functions swap elements as necessary from the two input sets
// to form a single output set, returning its size: the output set will be
// in the range [0:size], and will be in sorted order and free of
// duplicates. Elements which were moved into the range [size:Len] will have
// undefined order and may contain duplicates.
//
// All pivots must be in the range [0:Len]. A panic may occur when invalid
// pivots are passed into any of the functions.
//
// Convenience functions exist for slices of int, float64, and string
// element types, and also serve as examples for implementing utility
// functions for other types.
//
// Elements will be considered equal if `!Less(i,j) && !Less(j,i)`. An
// implication of this is that NaN values are equal to each other.
package set

import (
	"sort"

	"github.com/xtgo/set/internal"
	"github.com/xtgo/set/v2/setalgo"
)

// BUG(extemporalgenome): All ops should use binary search when runs are detected

// Op represents any of the mutating functions, such as Inter.
type Op func(data sort.Interface, pivot int) (size int)

// Cmp represents any of the comparison functions, such as IsInter.
type Cmp func(data sort.Interface, pivot int) bool

// Uniq swaps away duplicate elements in data, returning the size of the
// unique set. data is expected to be pre-sorted, and the resulting set in
// the range [0:size] will remain in sorted order. Uniq, following a
// sort.Sort call, can be used to prepare arbitrary inputs for use as sets.
func Uniq(data sort.Interface) (size int) {
	return setalgo.Uniq(data)
}

// Inter performs an in-place intersection on the two sets [0:pivot] and
// [pivot:Len]; the resulting set will occupy [0:size].
func Inter(data sort.Interface, pivot int) (size int) {
	return setalgo.Inter(data, pivot)
}

// Union performs an in-place union on the two sets [0:pivot] and
// [pivot:Len]; the resulting set will occupy [0:size].
func Union(data sort.Interface, pivot int) (size int) {
	return setalgo.Union(data, pivot)
}

// Diff performs an in-place difference on the two sets represented by
// [0:pivot] and [pivot:Len]; the resulting set will occupy [0:size].
//
// Do not pass Diff to the Apply function (Apply requires associative
// operations, which Diff is not).
func Diff(data sort.Interface, pivot int) (size int) {
	return setalgo.Diff(data, pivot)
}

// SymDiff performs an in-place symmetric difference on the two sets
// [0:pivot] and [pivot:Len]; the resulting set will occupy [0:size].
func SymDiff(data sort.Interface, pivot int) (size int) {
	return setalgo.SymDiff(data, pivot)
}

// IsSub returns true only if all elements in the range [0:pivot] are
// also present in the range [pivot:Len].
func IsSub(data sort.Interface, pivot int) bool {
	return setalgo.IsSub(data, pivot)
}

// IsSuper returns true only if all elements in the range [pivot:Len] are
// also present in the range [0:pivot]. IsSuper is especially useful for
// full membership testing.
func IsSuper(data sort.Interface, pivot int) bool {
	return setalgo.IsSuper(data, pivot)
}

// IsInter returns true if any element in the range [0:pivot] is also
// present in the range [pivot:Len]. IsInter is especially useful for
// partial membership testing.
func IsInter(data sort.Interface, pivot int) bool {
	return setalgo.IsInter(data, pivot)
}

// IsEqual returns true if the sets [0:pivot] and [pivot:Len] are equal.
func IsEqual(data sort.Interface, pivot int) bool {
	return setalgo.IsEqual(data, pivot)
}

// Ints sorts and deduplicates a slice of ints in place, returning the
// resulting set.
func Ints(data []int) []int {
	sort.Ints(data)
	n := setalgo.Uniq(sort.IntSlice(data))
	return data[:n]
}

// Float64s sorts and deduplicates a slice of float64s in place, returning
// the resulting set.
func Float64s(data []float64) []float64 {
	sort.Float64s(data)
	n := setalgo.Uniq(sort.Float64Slice(data))
	return data[:n]
}

// Strings sorts and deduplicates a slice of strings in place, returning
// the resulting set.
func Strings(data []string) []string {
	sort.Strings(data)
	n := setalgo.Uniq(sort.StringSlice(data))
	return data[:n]
}

// IntsDo applies op to the int sets, s and t, returning the result.
// s and t must already be individually sorted and free of duplicates.
func IntsDo(op Op, s []int, t ...int) []int {
	data := sort.IntSlice(append(s, t...))
	n := op(data, len(s))
	return data[:n]
}

// Float64sDo applies op to the float64 sets, s and t, returning the result.
// s and t must already be individually sorted and free of duplicates.
func Float64sDo(op Op, s []float64, t ...float64) []float64 {
	data := sort.Float64Slice(append(s, t...))
	n := op(data, len(s))
	return data[:n]
}

// StringsDo applies op to the string sets, s and t, returning the result.
// s and t must already be individually sorted and free of duplicates.
func StringsDo(op Op, s []string, t ...string) []string {
	data := sort.StringSlice(append(s, t...))
	n := op(data, len(s))
	return data[:n]
}

// IntsChk compares s and t according to cmp.
func IntsChk(cmp Cmp, s []int, t ...int) bool {
	data := sort.IntSlice(append(s, t...))
	return cmp(data, len(s))
}

// Float64sChk compares s and t according to cmp.
func Float64sChk(cmp Cmp, s []float64, t ...float64) bool {
	data := sort.Float64Slice(append(s, t...))
	return cmp(data, len(s))
}

// StringsChk compares s and t according to cmp.
func StringsChk(cmp Cmp, s []string, t ...string) bool {
	data := sort.StringSlice(append(s, t...))
	return cmp(data, len(s))
}

// Pivots transforms set-relative sizes into data-absolute pivots. Pivots is
// mostly only useful in conjunction with Apply. The sizes slice will be
// modified in-place and returned.
func Pivots(sizes ...int) []int {
	n := 0
	for i, l := range sizes {
		n += l
		sizes[i] = n
	}
	return sizes
}

// Apply concurrently applies op to all the sets terminated by pivots.
// pivots must contain one higher than the final index in each set, with the
// final element of pivots being equal to data.Len(); this deviates from the
// pivot semantics of other functions (which treat pivot as a delimiter) in
// order to make initializing the pivots slice simpler.
//
// data.Swap and data.Less are assumed to be concurrent-safe. Only
// associative operations should be used (Diff is not associative); see the
// Apply (Diff) example for a workaround. The result of applying SymDiff
// will contain elements that exist in an odd number of sets.
//
// The implementation runs op concurrently on pairs of neighbor sets
// in-place; when any pair has been merged, the resulting set is re-paired
// with one of its neighbor sets and the process repeats until only one set
// remains. The process is adaptive (large sets will not prevent small pairs
// from being processed), and strives for data-locality (only adjacent
// neighbors are paired and data shifts toward the zero index).
func Apply(op Op, data sort.Interface, pivots []int) (size int) {
	switch len(pivots) {
	case 0:
		return 0
	case 1:
		return pivots[0]
	case 2:
		return op(data, pivots[0])
	}

	spans := make([]internal.Span, 0, len(pivots)+1)

	// convert pivots into spans (index intervals that represent sets)
	i := 0
	for _, j := range pivots {
		spans = append(spans, internal.Span{i, j})
		i = j
	}

	n := len(spans) // original number of spans
	m := n / 2      // original number of span pairs (rounded down)

	// true if the span is being used
	inuse := make([]bool, n)

	ch := make(chan internal.Span, m)

	// reverse iterate over every other span, starting with the last;
	// concurrent algo (further below) will pick available pairs operate on
	for i := range spans[:m] {
		i = len(spans) - 1 - i*2
		ch <- spans[i]
	}

	for s := range ch {
		if len(spans) == 1 {
			if s.I != 0 {
				panic("impossible final span")
			}
			// this was the last operation
			return s.J
		}

		// locate the span we received (match on start of span only)
		i := sort.Search(len(spans), func(i int) bool { return spans[i].I >= s.I })

		// store the result (this may change field j but not field i)
		spans[i] = s

		// mark the span as available for use
		inuse[i] = false

		// check the immediate neighbors for availability (prefer left)
		j, k := i-1, i+1
		switch {
		case j >= 0 && !inuse[j]:
			i, j = j, i
		case k < len(spans) && !inuse[k]:
			j = k
		default:
			// nothing to do right now. wait for something else to finish
			continue
		}

		s, t := spans[i], spans[j]

		go func(s, t internal.Span) {
			// sizes of the respective sets
			l, m := s.J-s.I, t.J-t.I

			// shift the right-hand span to be adjacent to the left
			internal.Slide(data, s.J, t.I, m)

			// prepare a view of the data (abs -> rel indices)
			b := internal.BoundSpan{data, internal.Span{s.I, s.J + m}}

			// store result of op, adjusting for view (rel -> abs)
			s.J = s.I + op(b, l)

			// send the result back to the coordinating goroutine
			ch <- s
		}(s, t)

		// account for the spawn merging that will occur
		s.J += t.J - t.I

		k = j + 1

		// shrink the lists to account for the merger
		spans = append(append(spans[:i], s), spans[k:]...)

		// (and the merged span is now in use as well)
		inuse = append(append(inuse[:i], true), inuse[k:]...)
	}
	panic("unreachable")
}
