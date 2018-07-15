// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package setalgo

import (
	"sort"

	"github.com/xtgo/set/internal"
)

// Uniq swaps away duplicate elements in data, returning the size of the
// unique set. data is expected to be pre-sorted, and the resulting set in
// the range [0:size] will remain in sorted order. Uniq, following a
// sort.Sort call, can be used to prepare arbitrary inputs for use as sets.
func Uniq(data sort.Interface) (size int) {
	p, l := 0, data.Len()
	if l <= 1 {
		return l
	}
	for i := 1; i < l; i++ {
		if !data.Less(p, i) {
			continue
		}
		p++
		if p < i {
			data.Swap(p, i)
		}
	}
	return p + 1
}

// Inter performs an in-place intersection on the two sets [0:pivot] and
// [pivot:Len]; the resulting set will occupy [0:size].
func Inter(data sort.Interface, pivot int) (size int) {
	k, l := pivot, data.Len()
	p, i, j := 0, 0, k
	for i < k && j < l {
		switch {
		case data.Less(i, j):
			i++
		case data.Less(j, i):
			j++
		case p < i:
			data.Swap(p, i)
			fallthrough
		default:
			p, i, j = p+1, i+1, j+1
		}
	}
	return p
}

// Union performs an in-place union on the two sets [0:pivot] and
// [pivot:Len]; the resulting set will occupy [0:size].
func Union(data sort.Interface, pivot int) (size int) {
	// BUG(extemporalgenome): Union currently uses a multi-pass implementation

	sort.Sort(data)
	return Uniq(data)

	flipped := false
	_ = flipped // TODO
	k, l := pivot, data.Len()
	p, i, j := 0, 0, k
	for i < k && j < l {
		switch {
		case data.Less(j, i):
			data.Swap(i, j)
		case data.Less(i, j):
			if p < i {
				data.Swap(p, i)
			}
			p, i = p+1, i+1
		}
	}
	panic("not implemented")
}

// Diff performs an in-place difference on the two sets represented by
// [0:pivot] and [pivot:Len]; the resulting set will occupy [0:size].
func Diff(data sort.Interface, pivot int) (size int) {
	k, l := pivot, data.Len()
	p, i, j := 0, 0, k
	for i < k && j < l {
		switch {
		case data.Less(i, j):
			if p < i {
				data.Swap(p, i)
			}
			p, i = p+1, i+1
		case data.Less(j, i):
			j++
		default:
			i, j = i+1, j+1
		}
	}
	// at this point we've exhausted one of the sets, so swap any
	// remaining elements, if any, from the first set into place.
	return internal.XSwap(data, p, i, k, k)
}

// SymDiff performs an in-place symmetric difference on the two sets
// [0:pivot] and [pivot:Len]; the resulting set will occupy [0:size].
func SymDiff(data sort.Interface, pivot int) (size int) {
	// BUG(extemporalgenome): SymDiff currently uses a multi-pass implementation

	i := Inter(data, pivot)
	l := data.Len()
	b := internal.BoundSpan{data, internal.Span{i, l}}
	sort.Sort(b)
	size = Uniq(b)
	internal.Slide(data, 0, i, size)
	l = i + size
	sort.Sort(internal.BoundSpan{data, internal.Span{size, l}})
	return Diff(data, size)
}
