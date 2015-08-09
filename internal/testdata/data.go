// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

import "reflect"

type BinTest struct {
	A, B    []int
	Inter   []int
	Union   []int
	Diff    []int
	RevDiff []int
	SymDiff []int
	IsSub   bool
	IsSuper bool
	IsInter bool
	IsEqual bool
}

func (t BinTest) sel(name string) interface{} { return reflect.ValueOf(t).FieldByName(name).Interface() }
func (t BinTest) SelSlice(name string) []int  { return t.sel(name).([]int) }
func (t BinTest) SelBool(name string) bool    { return t.sel(name).(bool) }

var BinTests = []BinTest{
	{
		// empty sets
		A:       nil,
		B:       nil,
		Inter:   nil,
		Union:   nil,
		Diff:    nil,
		RevDiff: nil,
		SymDiff: nil,
		IsSub:   true,
		IsSuper: true,
		IsInter: false,
		IsEqual: true,
	},
	{
		// identical sets
		A:       []int{1, 2, 3},
		B:       []int{1, 2, 3},
		Inter:   []int{1, 2, 3},
		Union:   []int{1, 2, 3},
		Diff:    nil,
		RevDiff: nil,
		SymDiff: nil,
		IsSub:   true,
		IsSuper: true,
		IsInter: true,
		IsEqual: true,
	},
	{
		// non-disjoint sets
		A:       []int{1, 2, 3},
		B:       []int{2, 3, 4},
		Inter:   []int{2, 3},
		Union:   []int{1, 2, 3, 4},
		Diff:    []int{1},
		RevDiff: []int{4},
		SymDiff: []int{1, 4},
		IsSub:   false,
		IsSuper: false,
		IsInter: true,
		IsEqual: false,
	},
	{
		// inverse non-disjoint sets
		A:       []int{2, 3, 4},
		B:       []int{1, 2, 3},
		Inter:   []int{2, 3},
		Union:   []int{1, 2, 3, 4},
		Diff:    []int{4},
		RevDiff: []int{1},
		SymDiff: []int{1, 4},
		IsSub:   false,
		IsSuper: false,
		IsInter: true,
		IsEqual: false,
	},
	{
		// disjoint sets
		A:       []int{1, 2, 3},
		B:       []int{4, 5, 6},
		Inter:   nil,
		Union:   []int{1, 2, 3, 4, 5, 6},
		Diff:    []int{1, 2, 3},
		RevDiff: []int{4, 5, 6},
		SymDiff: []int{1, 2, 3, 4, 5, 6},
		IsSub:   false,
		IsSuper: false,
		IsInter: false,
		IsEqual: false,
	},
	{
		// inverse disjoint sets
		A:       []int{4, 5, 6},
		B:       []int{1, 2, 3},
		Inter:   nil,
		Union:   []int{1, 2, 3, 4, 5, 6},
		Diff:    []int{4, 5, 6},
		RevDiff: []int{1, 2, 3},
		SymDiff: []int{1, 2, 3, 4, 5, 6},
		IsSub:   false,
		IsSuper: false,
		IsInter: false,
		IsEqual: false,
	},
	{
		// alternating disjoint sets
		A:       []int{1, 3, 5},
		B:       []int{2, 4, 6},
		Inter:   nil,
		Union:   []int{1, 2, 3, 4, 5, 6},
		Diff:    []int{1, 3, 5},
		RevDiff: []int{2, 4, 6},
		SymDiff: []int{1, 2, 3, 4, 5, 6},
		IsSub:   false,
		IsSuper: false,
		IsInter: false,
		IsEqual: false,
	},
	{
		// inverse alternating disjoint sets
		A:       []int{2, 4, 6},
		B:       []int{1, 3, 5},
		Inter:   nil,
		Union:   []int{1, 2, 3, 4, 5, 6},
		Diff:    []int{2, 4, 6},
		RevDiff: []int{1, 3, 5},
		SymDiff: []int{1, 2, 3, 4, 5, 6},
		IsSub:   false,
		IsSuper: false,
		IsInter: false,
		IsEqual: false,
	},
	{
		// subset
		A:       []int{2},
		B:       []int{1, 2, 3},
		Inter:   []int{2},
		Union:   []int{1, 2, 3},
		Diff:    nil,
		RevDiff: []int{1, 3},
		SymDiff: []int{1, 3},
		IsSub:   true,
		IsSuper: false,
		IsInter: true,
		IsEqual: false,
	},
	{
		// superset
		A:       []int{1, 2, 3},
		B:       []int{2},
		Inter:   []int{2},
		Union:   []int{1, 2, 3},
		Diff:    []int{1, 3},
		RevDiff: nil,
		SymDiff: []int{1, 3},
		IsSub:   false,
		IsSuper: true,
		IsInter: true,
		IsEqual: false,
	},
}
