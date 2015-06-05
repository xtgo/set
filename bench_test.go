// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set_test

import (
	"sort"
	"testing"

	"github.com/xtgo/set"
)

func seq(start, stop, skip int) []int {
	n := (stop - start) / skip
	s := make([]int, n)
	for i := range s {
		s[i] = (start + i) * skip
	}
	return s
}

func interleave(n, k int) [][]int {
	l := n * k
	sets := make([][]int, n)
	for i := range sets {
		sets[i] = seq(i, i+l, n)
	}
	return sets
}

func concat(n, k, gap int) [][]int {
	l := k + gap
	sets := make([][]int, n)
	for i := range sets {
		start := i * l
		sets[i] = seq(start, start+l, 1)
	}
	return sets
}

func reverse(sets [][]int) [][]int {
	n := len(sets)
	for i := range sets[:n/2] {
		j := n - i - 1
		sets[i], sets[j] = sets[j], sets[i]
	}
	return sets
}

func benchOp(b *testing.B, op set.Op, sets [][]int) {
	s, t := sets[0], sets[1]
	data := make(sort.IntSlice, 0, len(s)+len(t))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//b.StopTimer()
		data = append(append(data[:0], s...), t...)
		//b.StartTimer()
		op(data, len(s))
	}
}

func benchMapOp(b *testing.B, op func(s, t IntSet), sets [][]int) {
	s, t := NewIntSet(sets[0]), NewIntSet(sets[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//b.StopTimer()
		x, y := CopyIntSet(s), CopyIntSet(t)
		//b.StartTimer()
		op(x, y)
	}
}

const (
	small = 32
	large = 64 * 1024
)

// union, inter: requires most Swap calls, fewest Less calls
func dataRevCat(n int, size int) [][]int { return reverse(concat(n, size, 0)) }

// union, inter: requires ~most Swap calls, most Less calls
func dataAlternate(n int, size int) [][]int { return interleave(n, size) }

func dataOverlap(n int, size int) [][]int { return concat(n, size, -size/2) }

func BenchmarkUnion64K_revcat(b *testing.B)    { benchOp(b, set.Union, dataRevCat(2, large)) }
func BenchmarkMapUnion64K_revcat(b *testing.B) { benchMapOp(b, IntSet.Union, dataRevCat(2, large)) }

func BenchmarkUnion32(b *testing.B)     { benchOp(b, set.Union, dataOverlap(2, small)) }
func BenchmarkMapUnion32(b *testing.B)  { benchMapOp(b, IntSet.Union, dataOverlap(2, small)) }
func BenchmarkUnion64K(b *testing.B)    { benchOp(b, set.Union, dataOverlap(2, large)) }
func BenchmarkMapUnion64K(b *testing.B) { benchMapOp(b, IntSet.Union, dataOverlap(2, large)) }

func BenchmarkUnion_alt32(b *testing.B)     { benchOp(b, set.Union, dataAlternate(2, small)) }
func BenchmarkMapUnion_alt32(b *testing.B)  { benchMapOp(b, IntSet.Union, dataAlternate(2, small)) }
func BenchmarkUnion_alt64K(b *testing.B)    { benchOp(b, set.Union, dataAlternate(2, large)) }
func BenchmarkMapUnion_alt64K(b *testing.B) { benchMapOp(b, IntSet.Union, dataAlternate(2, large)) }

func BenchmarkInter32(b *testing.B)     { benchOp(b, set.Inter, dataOverlap(2, small)) }
func BenchmarkMapInter32(b *testing.B)  { benchMapOp(b, IntSet.Inter, dataOverlap(2, small)) }
func BenchmarkInter64K(b *testing.B)    { benchOp(b, set.Inter, dataOverlap(2, large)) }
func BenchmarkMapInter64K(b *testing.B) { benchMapOp(b, IntSet.Inter, dataOverlap(2, large)) }

func BenchmarkInter_alt32(b *testing.B)     { benchOp(b, set.Inter, dataAlternate(2, small)) }
func BenchmarkMapInter_alt32(b *testing.B)  { benchMapOp(b, IntSet.Inter, dataAlternate(2, small)) }
func BenchmarkInter_alt64K(b *testing.B)    { benchOp(b, set.Inter, dataAlternate(2, large)) }
func BenchmarkMapInter_alt64K(b *testing.B) { benchMapOp(b, IntSet.Inter, dataAlternate(2, large)) }

func BenchmarkDiff32(b *testing.B)     { benchOp(b, set.Diff, dataOverlap(2, small)) }
func BenchmarkMapDiff32(b *testing.B)  { benchMapOp(b, IntSet.Diff, dataOverlap(2, small)) }
func BenchmarkDiff64K(b *testing.B)    { benchOp(b, set.Diff, dataOverlap(2, large)) }
func BenchmarkMapDiff64K(b *testing.B) { benchMapOp(b, IntSet.Diff, dataOverlap(2, large)) }

func BenchmarkDiff_alt32(b *testing.B)     { benchOp(b, set.Diff, dataAlternate(2, small)) }
func BenchmarkMapDiff_alt32(b *testing.B)  { benchMapOp(b, IntSet.Diff, dataAlternate(2, small)) }
func BenchmarkDiff_alt64K(b *testing.B)    { benchOp(b, set.Diff, dataAlternate(2, large)) }
func BenchmarkMapDiff_alt64K(b *testing.B) { benchMapOp(b, IntSet.Diff, dataAlternate(2, large)) }

func BenchmarkSymDiff32(b *testing.B)     { benchOp(b, set.SymDiff, dataOverlap(2, small)) }
func BenchmarkMapSymDiff32(b *testing.B)  { benchMapOp(b, IntSet.SymDiff, dataOverlap(2, small)) }
func BenchmarkSymDiff64K(b *testing.B)    { benchOp(b, set.SymDiff, dataOverlap(2, large)) }
func BenchmarkMapSymDiff64K(b *testing.B) { benchMapOp(b, IntSet.SymDiff, dataOverlap(2, large)) }

func BenchmarkSymDiff_alt32(b *testing.B)     { benchOp(b, set.SymDiff, dataAlternate(2, small)) }
func BenchmarkMapSymDiff_alt32(b *testing.B)  { benchMapOp(b, IntSet.SymDiff, dataAlternate(2, small)) }
func BenchmarkSymDiff_alt64K(b *testing.B)    { benchOp(b, set.SymDiff, dataAlternate(2, large)) }
func BenchmarkMapSymDiff_alt64K(b *testing.B) { benchMapOp(b, IntSet.SymDiff, dataAlternate(2, large)) }
