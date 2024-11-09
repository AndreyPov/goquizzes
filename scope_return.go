package main

func f(n int) (r int) {
	a, r := n-1, n+1
	if a+a == r {
		c, r := n, n*n // r is redeclared in this block, it's not changing r from the outer block
		r = r - c
		/* The block could be rewritten and it would change r from outer scope:
		c := n
		r = n*n
		r = r - c
		*/
	}
	return
}

func main() {
	println(f(3)) // 4
}
