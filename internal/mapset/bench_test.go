// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapset_test

import (
	"testing"

	"github.com/xtgo/set/internal/mapset"
	td "github.com/xtgo/set/internal/testdata"
)

func BenchmarkUnion64K_revcat(b *testing.B) { benchOp(b, "Union", td.RevCat(2, td.Large)) }
func BenchmarkUnion32(b *testing.B)         { benchOp(b, "Union", td.Overlap(2, td.Small)) }
func BenchmarkUnion64K(b *testing.B)        { benchOp(b, "Union", td.Overlap(2, td.Large)) }
func BenchmarkUnion_alt32(b *testing.B)     { benchOp(b, "Union", td.Alternate(2, td.Small)) }
func BenchmarkUnion_alt64K(b *testing.B)    { benchOp(b, "Union", td.Alternate(2, td.Large)) }
func BenchmarkInter32(b *testing.B)         { benchOp(b, "Inter", td.Overlap(2, td.Small)) }
func BenchmarkInter64K(b *testing.B)        { benchOp(b, "Inter", td.Overlap(2, td.Large)) }
func BenchmarkInter_alt32(b *testing.B)     { benchOp(b, "Inter", td.Alternate(2, td.Small)) }
func BenchmarkInter_alt64K(b *testing.B)    { benchOp(b, "Inter", td.Alternate(2, td.Large)) }
func BenchmarkDiff32(b *testing.B)          { benchOp(b, "Diff", td.Overlap(2, td.Small)) }
func BenchmarkDiff64K(b *testing.B)         { benchOp(b, "Diff", td.Overlap(2, td.Large)) }
func BenchmarkDiff_alt32(b *testing.B)      { benchOp(b, "Diff", td.Alternate(2, td.Small)) }
func BenchmarkDiff_alt64K(b *testing.B)     { benchOp(b, "Diff", td.Alternate(2, td.Large)) }
func BenchmarkSymDiff32(b *testing.B)       { benchOp(b, "SymDiff", td.Overlap(2, td.Small)) }
func BenchmarkSymDiff64K(b *testing.B)      { benchOp(b, "SymDiff", td.Overlap(2, td.Large)) }
func BenchmarkSymDiff_alt32(b *testing.B)   { benchOp(b, "SymDiff", td.Alternate(2, td.Small)) }
func BenchmarkSymDiff_alt64K(b *testing.B)  { benchOp(b, "SymDiff", td.Alternate(2, td.Large)) }

func benchOp(b *testing.B, name string, sets [][]int) {
	var op mutOp
	td.ConvMethod(&op, mapset.Set(nil), name)
	s, t := mapset.New(sets[0]), mapset.New(sets[1])

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, y := s.Copy(), t.Copy()
		op(x, y)
	}
}
