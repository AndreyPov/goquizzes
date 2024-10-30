package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name
	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

// these 2 pointers are different because the pointer in the function
// is a copy of the original pointer, and the copy has a different address
// than the original pointer.
