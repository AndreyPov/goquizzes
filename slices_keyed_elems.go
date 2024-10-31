package main

import "fmt"

var x = []int{2: 5, 6, 0: 7}

func main() {
	// the index of an un-keyed element is determined by the index of the last keyed element.

	fmt.Println(x)
	// 7 - is the 0th element, 5 - is the 2nd element,
	// 6 - is the 3rd element (depends on the position of 5)

	var y = []int{4: 44, 55, 66, 1: 77, 88}
	fmt.Println(y)
	// 77 - is the 1st element, 88 - is the 2th element,
	// 44 - is the 4th element, 55 - is the 5th element, 66 - is the 6th element
}
