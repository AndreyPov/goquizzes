package main

func main() {
	count := 0
	for i := range [256]struct{}{} {
		if n := byte(i); n == -n { // byte is uint8
			count++
		}
	}
	println(count)
}
