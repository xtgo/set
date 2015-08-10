// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapset_test

import (
	"testing"

	"github.com/xtgo/set/internal/mapset"
	"github.com/xtgo/set/internal/testdata"
)

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
	mutOp  func(a, b mapset.Set) mapset.Set
	boolOp func(a, b mapset.Set) bool
)

func testMut(t *testing.T, name string) {
	var op mutOp
	testdata.ConvMethod(&op, mapset.Set(nil), name)

	for _, tt := range testdata.BinTests {
		a, b := mapset.New(tt.A), mapset.New(tt.B)
		c := op(a, b).Elems()
		want := tt.SelSlice(name)

		if !testdata.IsEqual(c, want) {
			t.Errorf(format, name, tt.A, tt.B, c, want)
		}
	}
}

func testBool(t *testing.T, name string) {
	var op boolOp
	testdata.ConvMethod(&op, mapset.Set(nil), name)

	for _, tt := range testdata.BinTests {
		a, b := mapset.New(tt.A), mapset.New(tt.B)
		ok := op(a, b)
		want := tt.SelBool(name)

		if ok != want {
			t.Errorf(format, name, tt.A, tt.B, ok, want)
		}
	}
}
