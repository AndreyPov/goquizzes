package main

func main() {
	i, s := 9, []int{}

	for i = range s {
	}
	print(i) // 9 because s is nil

	for i = 0; i < len(s); i++ {
	}
	print(i) // 0 because len(s) is 0 and the loop never runs but i is still assigned

	s = append(s, 1, 2, 3, 4, 5)

	for i = range s {
	}
	print(i) // 4 because the loop runs 5 times and i is assigned to the last index

	for i = 0; i < len(s); i++ {
	}
	println(i) // 5 because the loop runs 5 times and i is assigned to the last index
}
