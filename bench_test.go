// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set_test

import (
	"sort"
	"testing"

	"github.com/xtgo/set"
	"github.com/xtgo/set/internal/sliceset"
	td "github.com/xtgo/set/internal/testdata"
)

func BenchmarkUnion64K_revcat(b *testing.B) { benchMut(b, "Union", td.RevCat(2, td.Large)) }
func BenchmarkUnion32(b *testing.B)         { benchMut(b, "Union", td.Overlap(2, td.Small)) }
func BenchmarkUnion64K(b *testing.B)        { benchMut(b, "Union", td.Overlap(2, td.Large)) }
func BenchmarkUnion_alt32(b *testing.B)     { benchMut(b, "Union", td.Alternate(2, td.Small)) }
func BenchmarkUnion_alt64K(b *testing.B)    { benchMut(b, "Union", td.Alternate(2, td.Large)) }
func BenchmarkInter32(b *testing.B)         { benchMut(b, "Inter", td.Overlap(2, td.Small)) }
func BenchmarkInter64K(b *testing.B)        { benchMut(b, "Inter", td.Overlap(2, td.Large)) }
func BenchmarkInter_alt32(b *testing.B)     { benchMut(b, "Inter", td.Alternate(2, td.Small)) }
func BenchmarkInter_alt64K(b *testing.B)    { benchMut(b, "Inter", td.Alternate(2, td.Large)) }
func BenchmarkDiff32(b *testing.B)          { benchMut(b, "Diff", td.Overlap(2, td.Small)) }
func BenchmarkDiff64K(b *testing.B)         { benchMut(b, "Diff", td.Overlap(2, td.Large)) }
func BenchmarkDiff_alt32(b *testing.B)      { benchMut(b, "Diff", td.Alternate(2, td.Small)) }
func BenchmarkDiff_alt64K(b *testing.B)     { benchMut(b, "Diff", td.Alternate(2, td.Large)) }
func BenchmarkSymDiff32(b *testing.B)       { benchMut(b, "SymDiff", td.Overlap(2, td.Small)) }
func BenchmarkSymDiff64K(b *testing.B)      { benchMut(b, "SymDiff", td.Overlap(2, td.Large)) }
func BenchmarkSymDiff_alt32(b *testing.B)   { benchMut(b, "SymDiff", td.Alternate(2, td.Small)) }
func BenchmarkSymDiff_alt64K(b *testing.B)  { benchMut(b, "SymDiff", td.Alternate(2, td.Large)) }

func BenchmarkIsInter32(b *testing.B)      { benchBool(b, "IsInter", td.Overlap(2, td.Small)) }
func BenchmarkIsInter64K(b *testing.B)     { benchBool(b, "IsInter", td.Overlap(2, td.Large)) }
func BenchmarkIsInter_alt32(b *testing.B)  { benchBool(b, "IsInter", td.Alternate(2, td.Small)) }
func BenchmarkIsInter_alt64K(b *testing.B) { benchBool(b, "IsInter", td.Alternate(2, td.Large)) }

func BenchmarkApply256_64K(b *testing.B) { benchApply(b, td.Rand(256, td.Large)) }

func benchMut(b *testing.B, name string, sets [][]int) {
	var op mutOp
	td.ConvMethod(&op, sliceset.Set(nil), name)
	bench(b, func(a, b sliceset.Set) { op(a, b) }, sets)
}

func benchBool(b *testing.B, name string, sets [][]int) {
	var op boolOp
	td.ConvMethod(&op, sliceset.Set(nil), name)
	bench(b, func(a, b sliceset.Set) { op(a, b) }, sets)
}

func bench(b *testing.B, op func(a, b sliceset.Set), sets [][]int) {
	s, t := sets[0], sets[1]
	data := make([]int, 0, len(s)+len(t))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data = append(data[:0], s...)
		op(data, t)
	}
}

func pivots(sets [][]int) []int {
	lengths := make([]int, len(sets))
	for i, set := range sets {
		lengths[i] = len(set)
	}
	return set.Pivots(lengths...)
}

func benchApply(b *testing.B, sets [][]int) {
	pivots := pivots(sets)
	n := len(sets) - 1
	data := make(sort.IntSlice, 0, pivots[n])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data = data[:0]
		for _, set := range sets {
			data = append(data, set...)
		}
		set.Apply(set.Inter, data, pivots)
	}
}
