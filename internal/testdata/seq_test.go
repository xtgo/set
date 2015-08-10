package testdata

import "fmt"

func ExampleAlternate() {
	fmt.Println(Alternate(2, 4))
	// Output: [[0 2 4 6] [1 3 5 7]]
}

func ExampleOverlap() {
	fmt.Println(Overlap(2, 4))
	// Output: [[0 1 2 3] [2 3 4 5]]
}
