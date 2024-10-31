package main

func main() {
	x := []int{7, 8, 9}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			print(v)
		}()
		y[i] = &i
	}
	print(*y[0], *y[1], *y[2], " ")

	// Output: 222 999 (before GO 1.22), 012 987 (since Go 1.22)

	/* Explanation:
	- Prior to Go 1.22, the iteration variables (here they are i and v)
	are shared by all loop steps. In other words, each iteration variable only
	has one instance during the execution of the loop. That means all the elements
	of y store the same address (of the value v). In the end, the variable i is set as 2,
	and the variable v is set as 9.
	So, prior to Go 1.22, the output is 222 999.

	- Since Go 1.22, the iteration variables (here they are i and v) will get new instances
	at each loop step. So, since Go 1.22, the output is 012 987.
	*/
}
