package main

var xx = *new(*int)
var y *int = nil

func fer() interface{} {
	return y
}

func main() {
	if fer() == nil {
		if xx == nil {
			println("A")
		} else {
			println("B")
		}
	} else {
		if xx == nil {
			println("C")
		} else {
			println("D")
		}
	}
}

/* Output: C
Explanation:
An interface value with a nil non-interface dynamic value is not a nil interface.
So the function f always return a non-nil interface value.
A call to the new builtin function always returns a pointer pointing to a zero value.
That means dereferencing the pointer results in a zero value.
The zero value of the *int type is nil.
*/
