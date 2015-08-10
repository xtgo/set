// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testdata

import "math/rand"

const (
	Small = 32
	Large = 64 * 1024
)

func Seq(start, stop, skip int) []int {
	n := (stop - start) / skip
	s := make([]int, n)
	for i := range s {
		s[i] = start + (i * skip)
	}
	return s
}

func Interleave(n, k int) [][]int {
	l := n * k
	sets := make([][]int, n)
	for i := range sets {
		sets[i] = Seq(i, i+l, n)
	}
	return sets
}

func Concat(n, k, gap int) [][]int {
	l := k + gap
	sets := make([][]int, n)
	for i := range sets {
		start := i * l
		sets[i] = Seq(start, start+k, 1)
	}
	return sets
}

func Reverse(sets [][]int) [][]int {
	n := len(sets)
	for i := range sets[:n/2] {
		j := n - i - 1
		sets[i], sets[j] = sets[j], sets[i]
	}
	return sets
}

func RevCat(n int, size int) [][]int {
	// union, inter: requires most Swap calls, fewest Less calls
	return Reverse(Concat(n, size, 0))
}

func Alternate(n int, size int) [][]int {
	// union, inter: requires ~most Swap calls, most Less calls
	return Interleave(n, size)
}

func Overlap(n int, size int) [][]int {
	return Concat(n, size, -size/2)
}

func Rand(n int, size int) [][]int {
	rand.Seed(0)
	sets := make([][]int, n)
	for i := range sets {
		start, l := rand.Intn(size), rand.Intn(size)+1
		stop := start + l
		sets[i] = Seq(start, stop, 1)
	}
	return sets
}
