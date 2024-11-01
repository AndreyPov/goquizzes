package main

func f() bool {
	return false
}

func main() {
	switch f(); //
	{
	case true:
		println(1)
	case false:
		println(0)
	}
}

// Output: 1
/* Explanation:
semicolons are automatically inserted by the compiler, so the code is equivalent to:
switch f(); true {
case true:
	println(1)
case false:
	println(0)
}

true is the default value for a boolean in switch statements.
*/
