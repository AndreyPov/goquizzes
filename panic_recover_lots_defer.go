package main

func main() {
	defer func() {
		println(recover().(int))
	}()
	defer func() {
		defer func() {
			recover()
		}()
		defer recover()
		panic(3)
	}()
	defer func() {
		defer func() {
			defer func() {
				recover()
			}()
		}()
		defer recover()
		panic(2)
	}()
	panic(1)
}

// Output: 2

/* Explanation:
- Except the two recover calls at line 9 and 5, the other ones are all no-op.
- The recover calls at line 9 recovers the panic 3.
- The recover calls at line 5 recovers the panic 2.
- The the panic 1 is never recovered, but it is suppressed by the panic 2.
*/

// Link with explanation: https://go101.org/quizzes/panic-recover-2.html#:~:text=Explain%20Panic/Recover%20Mechanism%20in%20Detail
