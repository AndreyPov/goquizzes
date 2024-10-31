package main

func main() {
	c := make(chan int, 1)
	for done := false; !done; {
		select {
		default:
			print(1)
			done = true
		case <-c:
			print(2)
			c = nil // this action makes the channel c unreachable,
			// without that the program will loop forever
		case c <- 1:
			print(3)
		}
	}
	// Output: 321, as c <- 1 is executed first, then <-c is executed, and then default is executed
}
