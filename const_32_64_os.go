package main

const X = '\x61' // 'a' in ASCII, it's rune or int32
const Y = 0x62   // int constant, here it is int64, because OS is 64-bit
const N = Y - X  // 1
const M int64 = 1

var n = 32

func main() {
	if N == M {
		println(X, Y, N, M, N<<n>>n, M<<n>>n)
	}
}
