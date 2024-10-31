package main

import "fmt"

func main() {
	m := make(map[string]int, 3)
	// maps have no capacity, len(m) is 0 here, 3 is just a hint

	x := len(m)
	m["Go"] = m["Go"] // which is 0 in fact
	y := len(m)       // len becomes 1

	fmt.Println(x, y)
}
