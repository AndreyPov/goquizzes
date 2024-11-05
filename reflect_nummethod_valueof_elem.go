package main

import (
	"fmt"
	"reflect"
)

type I interface {
	foo()
	Bar()
}

type T struct {
	a string
}

func (T) foo() {}
func (T) Bar() {}

func main() {
	var t T = T{"sdfg"}
	var i I = t
	var x = reflect.ValueOf(t)
	var y = reflect.ValueOf(&i).Elem()
	fmt.Println(x, y)                         // {sdfg} {sdfg} - means just returns the value of the interface
	fmt.Println(x.NumMethod(), y.NumMethod()) // Output: 1 2

	/* Explanation:
	NumMethod returns the number of methods in the value's method set.
	For a non-interface type, it returns the number of exported methods.
	For an interface type, it returns the number of exported and unexported methods.
	*/
}
