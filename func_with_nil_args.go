package main

func farg(vs ...interface{}) {
	print(len(vs))
}

func main() {
	farg()       // exquivalent to farg(nil...)
	farg(nil)    // f(nil) is treated as an element in a slice. f(nil) is the same as f([]interface{}{nil}...)
	farg(nil...) // nil slice, or []interface{}
}
