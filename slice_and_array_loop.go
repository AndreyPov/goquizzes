package main

func c() {
	var a = [2]int{5, 7}
	for i, v := range a { // a is an array
		if i == 0 {
			a[1] = 9
		} else {
			print(v)
		}
	}
}

func g() {
	var a = [2]int{5, 7}
	for i, v := range a[:] { // a[:] is a slice
		if i == 0 {
			a[1] = 9
		} else {
			print(v)
		}
	}
}

func main() {
	c()
	g()
	// Output 79, when range on an array, Go takes a copy of the array
	// When range iterates over a slice, it refers to the actual underlying array.
	// Therefore, any modifications to a inside the loop are visible in subsequent iterations.
}
