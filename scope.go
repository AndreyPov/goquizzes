package main

var anotherF = func(x int) {}

func Bar() {
	anotherF := func(x int) {
		if x >= 0 {
			print(x)
			anotherF(x - 1)
		}
	}
	anotherF(2)
}

func Fooo() {
	anotherF = func(x int) {
		if x >= 0 {
			print(x)
			anotherF(x - 1)
		}
	}
	anotherF(2)
}

func main() {
	Bar()
	print(" | ")
	Fooo()
}
