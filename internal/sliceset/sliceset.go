// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sliceset provides a convenient []int set wrapper to aid in
// testing and benchmarks, and to serve as an example for those in need of
// a (concrete) abstraction for simplifying code. It is not intended for
// direct reuse.
package sliceset

import (
	"sort"

	"github.com/xtgo/set"
)

type Set []int

func (s Set) Len() int           { return len(s) }
func (s Set) Less(i, j int) bool { return s[i] < s[j] }
func (s Set) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s Set) Copy() Set { return append(Set(nil), s...) }

func (s Set) Union(t Set) Set    { return s.Do(set.Union, t) }
func (s Set) Inter(t Set) Set    { return s.Do(set.Inter, t) }
func (s Set) Diff(t Set) Set     { return s.Do(set.Diff, t) }
func (s Set) SymDiff(t Set) Set  { return s.Do(set.SymDiff, t) }
func (s Set) IsSub(t Set) bool   { return s.DoBool(set.IsSub, t) }
func (s Set) IsSuper(t Set) bool { return s.DoBool(set.IsSuper, t) }
func (s Set) IsInter(t Set) bool { return s.DoBool(set.IsInter, t) }
func (s Set) IsEqual(t Set) bool { return s.DoBool(set.IsEqual, t) }

func (s Set) Uniq() Set {
	n := set.Uniq(s)
	return s[:n]
}

func (s Set) Do(op set.Op, t Set) Set {
	data := append(s, t...)
	n := op(data, len(s))
	return data[:n]
}

type BoolOp func(sort.Interface, int) bool

func (s Set) DoBool(op BoolOp, t Set) bool {
	data := append(s, t...)
	return op(data, len(s))
}
