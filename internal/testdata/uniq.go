// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

type UniqTest struct {
	In, Out []int
}

var UniqTests = []UniqTest{
	{
		In:  nil,
		Out: nil,
	},
	{
		In:  []int{0, 1, 2, 3, 4, 5},
		Out: []int{0, 1, 2, 3, 4, 5},
	},
	{
		In:  []int{0, 0, 1, 2, 3, 3},
		Out: []int{0, 1, 2, 3},
	},
	{
		In:  []int{0, 1, 1, 1, 1, 2},
		Out: []int{0, 1, 2},
	},
}
