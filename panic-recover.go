package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		defer panic(1)
		recover()
	}()
	defer recover()
	panic(2)
}

//Output: 21
/* Explanation:
- The two recover calls at line 14 and 12 are no-op.
- The recover calls at line 10 catches the panic 2.
- The recover calls at line 7 catches the panic 1.
*/
