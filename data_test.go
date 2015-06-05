// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set_test

type SetTest struct {
	a, b    string
	inter   string
	union   string
	diff    string
	revdiff string
	symdiff string
	issub   bool
	issuper bool
	isinter bool
	isequal bool
}

var tests = []SetTest{
	{
		// empty sets
		a:       "",
		b:       "",
		inter:   "",
		union:   "",
		diff:    "",
		revdiff: "",
		symdiff: "",
		issub:   true,
		issuper: true,
		isinter: false,
		isequal: true,
	},
	{
		// identical sets
		a:       "abc",
		b:       "abc",
		inter:   "abc",
		union:   "abc",
		diff:    "",
		revdiff: "",
		symdiff: "",
		issub:   true,
		issuper: true,
		isinter: true,
		isequal: true,
	},
	{
		// non-disjoint sets
		a:       "abc",
		b:       "bcd",
		inter:   "bc",
		union:   "abcd",
		diff:    "a",
		revdiff: "d",
		symdiff: "ad",
		issub:   false,
		issuper: false,
		isinter: true,
		isequal: false,
	},
	{
		// inverse non-disjoint sets
		a:       "bcd",
		b:       "abc",
		inter:   "bc",
		union:   "abcd",
		diff:    "d",
		revdiff: "a",
		symdiff: "ad",
		issub:   false,
		issuper: false,
		isinter: true,
		isequal: false,
	},
	{
		// disjoint sets
		a:       "abc",
		b:       "def",
		inter:   "",
		union:   "abcdef",
		diff:    "abc",
		revdiff: "def",
		symdiff: "abcdef",
		issub:   false,
		issuper: false,
		isinter: false,
		isequal: false,
	},
	{
		// inverse disjoint sets
		a:       "def",
		b:       "abc",
		inter:   "",
		union:   "abcdef",
		diff:    "def",
		revdiff: "abc",
		symdiff: "abcdef",
		issub:   false,
		issuper: false,
		isinter: false,
		isequal: false,
	},
	{
		// alternating disjoint sets
		a:       "ace",
		b:       "bdf",
		inter:   "",
		union:   "abcdef",
		diff:    "ace",
		revdiff: "bdf",
		symdiff: "abcdef",
		issub:   false,
		issuper: false,
		isinter: false,
		isequal: false,
	},
	{
		// inverse alternating disjoint sets
		a:       "bdf",
		b:       "ace",
		inter:   "",
		union:   "abcdef",
		diff:    "bdf",
		revdiff: "ace",
		symdiff: "abcdef",
		issub:   false,
		issuper: false,
		isinter: false,
		isequal: false,
	},
	{
		// subset
		a:       "b",
		b:       "abc",
		inter:   "b",
		union:   "abc",
		diff:    "",
		revdiff: "ac",
		symdiff: "ac",
		issub:   true,
		issuper: false,
		isinter: true,
		isequal: false,
	},
	{
		// superset
		a:       "abc",
		b:       "b",
		inter:   "b",
		union:   "abc",
		diff:    "ac",
		revdiff: "",
		symdiff: "ac",
		issub:   false,
		issuper: true,
		isinter: true,
		isequal: false,
	},
}

var uniqs = []struct{ in, out []int }{
	{
		in:  nil,
		out: nil,
	},
	{
		in:  []int{0, 1, 2, 3, 4, 5},
		out: []int{0, 1, 2, 3, 4, 5},
	},
	{
		in:  []int{0, 0, 1, 2, 3, 3},
		out: []int{0, 1, 2, 3},
	},
	{
		in:  []int{0, 1, 1, 1, 1, 2},
		out: []int{0, 1, 2},
	},
}
