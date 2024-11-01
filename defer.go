package main

type Foo struct {
	v int
}

func MakeFoo(n *int) Foo {
	print(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	print(*n)
}

func main() {
	var x = 1
	var p = &x
	defer MakeFoo(p).Bar(p)
	x = 2
	p = new(int)
	MakeFoo(p)
}

//Output: 102
/* Explanation:
- MakeFoo(p) is called immediately when the line with defer is reached, though Bar(p) is deferred.
- Then x is changed to 2, and p is changed to a new address.
- p(pointer is changed to 0 and passed to MakeFoo, so the output is 0)
- The deferred function Bar(p) is called at the end of the main function, so the output is 2.
*/
