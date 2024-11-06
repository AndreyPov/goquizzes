package main

const (
	_    = 6
	_, _ = iota, iota + 10 // is the same as below
	A, _ = iota, iota + 10 // start from first const in the block, so A = 1 here.
	_, _
	_, B
)

func main() {
	println(A, B) // 1 13
}
