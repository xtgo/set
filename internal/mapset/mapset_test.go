// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapset_test

import (
	"testing"

	"github.com/xtgo/set/internal/mapset"
	"github.com/xtgo/set/internal/testdata"
)

func TestUnion(t *testing.T)   { runMut(t, "Union") }
func TestInter(t *testing.T)   { runMut(t, "Inter") }
func TestDiff(t *testing.T)    { runMut(t, "Diff") }
func TestSymDiff(t *testing.T) { runMut(t, "SymDiff") }
func TestIsSub(t *testing.T)   { runBool(t, "IsSub") }
func TestIsSuper(t *testing.T) { runBool(t, "IsSuper") }
func TestIsInter(t *testing.T) { runBool(t, "IsInter") }
func TestIsEqual(t *testing.T) { runBool(t, "IsEqual") }

const format = "%s(%v, %v) = %v, want %v"

type (
	mutOp  func(a, b mapset.Set) mapset.Set
	boolOp func(a, b mapset.Set) bool
)

func runMut(t *testing.T, name string) {
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

func runBool(t *testing.T, name string) {
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
