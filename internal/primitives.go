// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import "sort"

type Span struct{ I, J int }

// BoundSpan provides a sort.Interface for a subsection of data.
type BoundSpan struct {
	Data sort.Interface
	Span
}

func (b BoundSpan) Len() int           { return b.J - b.I }
func (b BoundSpan) Less(i, j int) bool { return b.Data.Less(b.I+i, b.I+j) }
func (b BoundSpan) Swap(i, j int)      { b.Data.Swap(b.I+i, b.I+j) }

// XSwap exchanges elements in the ranges [i:k] and [j:l], maintaining
// order. If the ranges are of different size, the shorter range will
// determine the number of swaps. XSwap will return i+min(k-i,l-j), which is
// the index proceeding the last element copied in place of the 'i' range.
func XSwap(data sort.Interface, i, j, k, l int) int {
	for i < k && j < l {
		data.Swap(i, j)
		i++
		j++
	}
	return i
}

// Slide shifts n elements from i to j.
func Slide(data sort.Interface, i, j, n int) {
	XSwap(data, i, j, i+n, j+n)
}
