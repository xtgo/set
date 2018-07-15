// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package setalgo_test

import (
	"testing"

	ss "github.com/xtgo/set/internal/sliceset"
	td "github.com/xtgo/set/internal/testdata"
)

func BenchmarkUnion64K_revcat(b *testing.B) { benchMut(b, ss.Set.Union, td.RevCat(2, td.Large)) }
func BenchmarkUnion32(b *testing.B)         { benchMut(b, ss.Set.Union, td.Overlap(2, td.Small)) }
func BenchmarkUnion64K(b *testing.B)        { benchMut(b, ss.Set.Union, td.Overlap(2, td.Large)) }
func BenchmarkUnion_alt32(b *testing.B)     { benchMut(b, ss.Set.Union, td.Alternate(2, td.Small)) }
func BenchmarkUnion_alt64K(b *testing.B)    { benchMut(b, ss.Set.Union, td.Alternate(2, td.Large)) }
func BenchmarkInter32(b *testing.B)         { benchMut(b, ss.Set.Inter, td.Overlap(2, td.Small)) }
func BenchmarkInter64K(b *testing.B)        { benchMut(b, ss.Set.Inter, td.Overlap(2, td.Large)) }
func BenchmarkInter_alt32(b *testing.B)     { benchMut(b, ss.Set.Inter, td.Alternate(2, td.Small)) }
func BenchmarkInter_alt64K(b *testing.B)    { benchMut(b, ss.Set.Inter, td.Alternate(2, td.Large)) }
func BenchmarkDiff32(b *testing.B)          { benchMut(b, ss.Set.Diff, td.Overlap(2, td.Small)) }
func BenchmarkDiff64K(b *testing.B)         { benchMut(b, ss.Set.Diff, td.Overlap(2, td.Large)) }
func BenchmarkDiff_alt32(b *testing.B)      { benchMut(b, ss.Set.Diff, td.Alternate(2, td.Small)) }
func BenchmarkDiff_alt64K(b *testing.B)     { benchMut(b, ss.Set.Diff, td.Alternate(2, td.Large)) }
func BenchmarkSymDiff32(b *testing.B)       { benchMut(b, ss.Set.SymDiff, td.Overlap(2, td.Small)) }
func BenchmarkSymDiff64K(b *testing.B)      { benchMut(b, ss.Set.SymDiff, td.Overlap(2, td.Large)) }
func BenchmarkSymDiff_alt32(b *testing.B)   { benchMut(b, ss.Set.SymDiff, td.Alternate(2, td.Small)) }
func BenchmarkSymDiff_alt64K(b *testing.B)  { benchMut(b, ss.Set.SymDiff, td.Alternate(2, td.Large)) }

func BenchmarkIsInter32(b *testing.B)      { benchBool(b, ss.Set.IsInter, td.Overlap(2, td.Small)) }
func BenchmarkIsInter64K(b *testing.B)     { benchBool(b, ss.Set.IsInter, td.Overlap(2, td.Large)) }
func BenchmarkIsInter_alt32(b *testing.B)  { benchBool(b, ss.Set.IsInter, td.Alternate(2, td.Small)) }
func BenchmarkIsInter_alt64K(b *testing.B) { benchBool(b, ss.Set.IsInter, td.Alternate(2, td.Large)) }

func benchMut(b *testing.B, op mutOp, sets [][]int) {
	bench(b, func(a, b ss.Set) { op(a, b) }, sets)
}

func benchBool(b *testing.B, op boolOp, sets [][]int) {
	bench(b, func(a, b ss.Set) { op(a, b) }, sets)
}

func bench(b *testing.B, op func(a, b ss.Set), sets [][]int) {
	s, t := sets[0], sets[1]
	data := make([]int, 0, len(s)+len(t))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data = append(data[:0], s...)
		op(data, t)
	}
}
