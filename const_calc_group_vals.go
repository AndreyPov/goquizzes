package main

const Q = 3

func main() {
	const (
		iota = iota
		X
		Y
	)
	println(X, Y)

	const (
		Q = Q + Q
		T // it's the same as T = Q + Q
	)

	println(Q, T)
}
