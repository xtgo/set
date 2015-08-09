// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mapset provides a reasonable map-based set implementation for
// use in comparative benchmarks, and to check arbitrary fuzz outputs.
// mapset is not intended for reuse.
package mapset

import "sort"

func New(s []int) Set {
	t := make(Set, len(s))
	for _, v := range s {
		t[v] = struct{}{}
	}
	return t
}

type Set map[int]struct{}

func (s Set) Union(t Set) Set {
	for v := range t {
		s[v] = struct{}{}
	}
	return s
}

func (s Set) Inter(t Set) Set {
	for v := range s {
		_, ok := t[v]
		if !ok {
			delete(s, v)
		}
	}
	return s
}

func (s Set) Diff(t Set) Set {
	for v := range t {
		delete(s, v)
	}
	return s
}

func (s Set) SymDiff(t Set) Set {
	for v := range t {
		_, ok := s[v]
		if ok {
			delete(s, v)
		} else {
			s[v] = struct{}{}
		}
	}
	return s
}

func (s Set) IsSub(t Set) bool {
	if len(s) > len(t) {
		return false
	}
	for k := range s {
		_, ok := t[k]
		if !ok {
			return false
		}
	}
	return true
}

func (s Set) IsSuper(t Set) bool {
	return t.IsSub(s)
}

func (s Set) IsInter(t Set) bool {
	for k := range s {
		_, ok := t[k]
		if ok {
			return true
		}
	}
	return false
}

func (s Set) IsEqual(t Set) bool {
	if len(s) != len(t) {
		return false
	}
	return s.IsSub(t)
}

func (s Set) Elems() []int {
	t := make([]int, 0, len(s))
	for v := range s {
		t = append(t, v)
	}
	sort.Ints(t)
	return t
}

func (s Set) Copy() Set {
	t := make(Set, len(s))
	t.Union(s)
	return t
}
