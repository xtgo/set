// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains a reasonable map-based set implementation for use in
// comparative benchmarks.

package set_test

var mote struct{}

type IntSet map[int]struct{}

func (s IntSet) Union(t IntSet) {
	for k := range t {
		s[k] = mote
	}
}

func (s IntSet) Inter(t IntSet) {
	for k := range s {
		_, ok := t[k]
		if !ok {
			delete(s, k)
		}
	}
}

func (s IntSet) Diff(t IntSet) {
	if len(t) <= len(s) {
		for k := range t {
			delete(s, k)
		}
		return
	}
	for k := range s {
		_, ok := t[k]
		if ok {
			delete(s, k)
		}
	}
}

func (s IntSet) SymDiff(t IntSet) {
	u, v := s, t
	if len(t) < len(s) {
		u, v = v, u
	}
	for k := range u {
		_, ok := v[k]
		if ok {
			delete(s, k)
			delete(t, k)
		}
	}
	for k := range t {
		s[k] = mote
	}
}

func (s IntSet) Apply(op func(s, t IntSet), sets ...IntSet) {
	for _, t := range sets {
		op(s, t)
	}
}

func NewIntSet(s []int) IntSet {
	t := make(IntSet, len(s))
	for _, k := range s {
		t[k] = mote
	}
	return t
}

func CopyIntSet(s IntSet) IntSet {
	t := make(IntSet, len(s))
	t.Union(s)
	return t
}
