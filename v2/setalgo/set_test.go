// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package setalgo_test

import (
	"testing"

	ss "github.com/xtgo/set/internal/sliceset"
	td "github.com/xtgo/set/internal/testdata"
)

func TestUniq(t *testing.T) {
	for _, tt := range td.UniqTests {
		s := ss.Set(tt.In).Copy().Uniq()

		if !td.IsEqual(s, tt.Out) {
			t.Errorf("Uniq(%v) = %v, want %v", tt.In, s, tt.Out)
		}
	}
}

func TestUnion(t *testing.T)   { testMut(t, "Union") }
func TestInter(t *testing.T)   { testMut(t, "Inter") }
func TestDiff(t *testing.T)    { testMut(t, "Diff") }
func TestSymDiff(t *testing.T) { testMut(t, "SymDiff") }
func TestIsSub(t *testing.T)   { testBool(t, "IsSub") }
func TestIsSuper(t *testing.T) { testBool(t, "IsSuper") }
func TestIsInter(t *testing.T) { testBool(t, "IsInter") }
func TestIsEqual(t *testing.T) { testBool(t, "IsEqual") }

const format = "%s(%v, %v) = %v, want %v"

type (
	mutOp  func(a, b ss.Set) ss.Set
	boolOp func(a, b ss.Set) bool
)

func testMut(t *testing.T, name string) {
	var op mutOp
	td.ConvMethod(&op, ss.Set(nil), name)

	for _, tt := range td.BinTests {
		a := ss.Set(tt.A).Copy()
		c := op(a, tt.B)
		want := tt.SelSlice(name)

		if !td.IsEqual(c, want) {
			t.Errorf(format, name, tt.A, tt.B, c, want)
		}
	}
}

func testBool(t *testing.T, name string) {
	var op boolOp
	td.ConvMethod(&op, ss.Set(nil), name)

	for _, tt := range td.BinTests {
		ok := op(tt.A, tt.B)
		want := tt.SelBool(name)

		if ok != want {
			t.Errorf(format, name, tt.A, tt.B, ok, want)
		}
	}
}
