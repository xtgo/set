// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set_test

import (
	"fmt"
	"sort"

	"github.com/xtgo/set"
)

func Example() {
	s := set.Strings([]string{"alpha", "gamma", "alpha"})
	fmt.Println("set:", s)

	s = set.StringsDo(set.Union, s, "beta")
	fmt.Println("set + beta:", s)

	fmt.Println(s, "contains any [alpha delta]:",
		set.StringsChk(set.IsInter, s, "alpha", "delta"))

	fmt.Println(s, "contains all [alpha delta]:",
		set.StringsChk(set.IsSuper, s, "alpha", "delta"))

	// Output:
	// set: [alpha gamma]
	// set + beta: [alpha beta gamma]
	// [alpha beta gamma] contains any [alpha delta]: true
	// [alpha beta gamma] contains all [alpha delta]: false
}

func ExampleUniq() {
	data := sort.IntSlice{5, 7, 3, 3, 5}

	sort.Sort(data)     // sort the data first
	n := set.Uniq(data) // Uniq returns the size of the set
	data = data[:n]     // trim the duplicate elements

	fmt.Println(data)
	// Output: [3 5 7]
}

func ExampleInter() {
	data := sort.IntSlice{3, 5, 7} // create a set (it must be sorted)
	pivot := len(data)             // store the length of our first set

	data = append(data, 1, 3, 5)   // append a second set (which also must be sorted)
	size := set.Inter(data, pivot) // perform set intersection

	// trim data to contain just the result set
	data = data[:size]

	fmt.Println("inter:", data)

	// Output:
	// inter: [3 5]
}

func ExampleIsSuper() {
	data := sort.StringSlice{"b", "c", "d"} // create a set (it must be sorted)
	pivot := len(data)                      // store the length of our first set

	data = append(data, "c", "d")         // append a second set (which also must be sorted)
	contained := set.IsSuper(data, pivot) // check the first set is a superset of the second

	fmt.Printf("%v superset of %v = %t\n", data[:pivot], data[pivot:], contained)

	data = data[:pivot] // trim off the second set

	data = append(data, "s")             // append a new singleton set to compare against
	contained = set.IsSuper(data, pivot) // check for membership

	fmt.Printf("%v superset of %v = %t\n", data[:pivot], data[pivot:], contained)

	// Output:
	// [b c d] superset of [c d] = true
	// [b c d] superset of [s] = false
}

func ExampleIsInter() {
	data := sort.StringSlice{"b", "c", "d"} // create a set (it must be sorted)
	pivot := len(data)                      // store the length of our first set

	data = append(data, "d", "z")         // append a second set (which also must be sorted)
	contained := set.IsInter(data, pivot) // check the first set is a superset of the second

	fmt.Printf("%v intersects %v = %t\n", data[:pivot], data[pivot:], contained)

	data = data[:pivot] // trim off the second set

	data = append(data, "s")             // append a new singleton set to compare against
	contained = set.IsInter(data, pivot) // check for membership

	fmt.Printf("%v intersects %v = %t\n", data[:pivot], data[pivot:], contained)

	// Output:
	// [b c d] intersects [d z] = true
	// [b c d] intersects [s] = false
}
