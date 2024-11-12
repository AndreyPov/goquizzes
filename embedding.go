package main

type Aa struct {
	g int
}

func (Aa) m() int {
	return 1
}

type Bb int

func (Bb) g() {}

func (Bb) f() {}

type C struct {
	Aa
	Bb
}

func (C) m() int {
	return 9
}

func main() {
	var c interface{} = C{}
	_, bf := c.(interface{ f() })
	_, bg := c.(interface{ g() })
	i := c.(interface{ m() int })
	println(bf, bg, i.m())
	// Output: true false 9
}

/* Explanation:
Field C.A.g and method C.B.g collide, so they are both not promoted.
Method C.B.f gets promoted as C.f.
Method C.m overrides C.A.m.
*/
