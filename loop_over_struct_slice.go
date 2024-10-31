package main

type Book struct {
	Pages int
}

func j() int {
	var books = []Book{{555}}
	for _, book := range books { // book is a copy of the element in the slice
		book.Pages = 999
	}
	return books[0].Pages // 555 because book is a copy
}

func k() int {
	var books = []*Book{{555}} // pointer to Book
	for _, book := range books {
		book.Pages = 999 // this will change the value in the slice
	}
	return books[0].Pages // 999 because book is a pointer
}

func main() {
	println(j(), k())
}
