// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/xtgo/set"
	"github.com/xtgo/set/setutil"
)

const format = "%s(%v) = %v, want %v"

type (
	BoolSpec struct {
		Name string
		Op   func(sort.Interface, int) bool
		Sel  func(SetTest) bool
	}

	SliceSpec struct {
		Name string
		Op   func(sort.Interface, int) int
		Sel  func(SetTest) string
	}
)

func setup(a, b string) (s string, pivot int) {
	return a + b, len(a)
}

func runMutator(t *testing.T, spec SliceSpec) {
	for _, tt := range tests {
		in, pivot := setup(tt.a, tt.b)
		s := setutil.Letters(in)
		//iface := LoggingInterface{t, s, pivot}
		iface := s
		l := spec.Op(iface, pivot)
		want := spec.Sel(tt)
		if want != string(s[:l]) {
			in := setutil.Letters(in).Mark(pivot, -1, -1)
			s := s.Mark(l, -1, -1)
			want := setutil.Letters(want).Mark(len(want), -1, -1)
			t.Errorf(format, spec.Name, in, s, want)
		}
	}
}

func TestInter(t *testing.T) {
	runMutator(t, SliceSpec{
		"Inter",
		set.Inter,
		func(tt SetTest) string { return tt.inter },
	})
}

func TestDiff(t *testing.T) {
	runMutator(t, SliceSpec{
		"Diff",
		set.Diff,
		func(tt SetTest) string { return tt.diff },
	})
}

func TestSymDiff(t *testing.T) {
	runMutator(t, SliceSpec{
		"SymDiff",
		set.SymDiff,
		func(tt SetTest) string { return tt.symdiff },
	})
}

func runBoolean(t *testing.T, spec BoolSpec) {
	for _, tt := range tests {
		in, pivot := setup(tt.a, tt.b)
		s := setutil.Letters(in)
		ok := spec.Op(s, pivot)
		want := spec.Sel(tt)
		if ok != want {
			t.Errorf(format, spec.Name, s, ok, want)
		}
	}
}

func TestIsSub(t *testing.T) {
	runBoolean(t, BoolSpec{
		"IsSub",
		set.IsSub,
		func(tt SetTest) bool { return tt.issub },
	})
}

func TestIsSuper(t *testing.T) {
	runBoolean(t, BoolSpec{
		"IsSuper",
		set.IsSuper,
		func(tt SetTest) bool { return tt.issuper },
	})
}

func TestIsInter(t *testing.T) {
	runBoolean(t, BoolSpec{
		"IsInter",
		set.IsInter,
		func(tt SetTest) bool { return tt.isinter },
	})
}

func TestIsEqual(t *testing.T) {
	runBoolean(t, BoolSpec{
		"IsEqual",
		set.IsEqual,
		func(tt SetTest) bool { return tt.isequal },
	})
}

func TestUniq(t *testing.T) {
	for _, tt := range uniqs {
		// make copy of the input
		s := sort.IntSlice(append([]int(nil), tt.in...))
		size := set.Uniq(s)
		s = s[:size]
		if len(s) != len(tt.out) {
			t.Errorf(format, "Uniq", tt.in, s, tt.out)
			continue
		}
		for i := range s {
			if s[i] != tt.out[i] {
				t.Errorf(format, "Uniq", tt.in, s, tt.out)
				break
			}
		}
	}
}

type LoggingInterface struct {
	t     *testing.T
	s     setutil.Letters
	pivot int
}

func (s LoggingInterface) Len() int {
	v := s.s.Len()
	s.t.Logf(">>> Len() = %v", v)
	return v
}

func (s LoggingInterface) Less(i, j int) bool {
	v := s.s.Less(i, j)
	pre := s.s.Mark(s.pivot, i, j)
	s.t.Logf(">>> %s Less(%d,%d) = %v", pre, i, j, v)
	return v
}

func (s LoggingInterface) Swap(i, j int) {
	pre := s.s.Mark(s.pivot, i, j)
	s.s.Swap(i, j)
	post := s.s.Mark(s.pivot, i, j)
	s.t.Logf(">>> %s Swap(%d,%d) %s", pre, i, j, post)
}

func (s LoggingInterface) String() string {
	return fmt.Sprint(s.s)
}
