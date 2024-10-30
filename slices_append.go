package main

func main() {

	// example 1
	var x = []string{"A", "B", "C"}

	for i, s := range x {
		print(i, s, ",")
		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}

	print("\n")

	// example 2

	var q = []string{"A", "B", "C"}

	for i, s := range q {
		print(i, s, ",")
		q = append(q, "Z")
		q[i+1] = "Z"
	}

	print("\n")

	// example 3

	var t = []string{"A", "B", "C", "D"}
	var r = t[:3]

	for i, s := range r {
		print(i, s, ",")
		r = append(r, "Z")
		r[i+1] = "Z"
	}
}

/* Output 1: 0A, 1M, 2C
Explanation 1:
- Ranging over a container iterates the elements of a copy of the container.
The evaluation of the copy happens before executing the loop,
so the length and capacity of the copy are never changed.
- A slice references its elements on a backing array.
So a copy of a slice shares the same elements (and the backing array) with the slice.
- The assignment x[i+1] = "M" in the first loop step modifies
the second element of the initial slice x and the copy of the initial slice x.
- If the free element slots of the first argument slice of an append call
are insufficient to hold all the appended elements,
the append call will create a new backing array to hold all the elements
of the first argument slice and the appended elements.
So, at the end of the first loop step, the backing array of the slice x is changed.
However, the change doesn't affect the slice copy used in the element iteration.
All subsequent element modifications apply to the new backing array,
so they have no effects on the copy used in the element iteration.
*/

// also in the last example note that append changes the slice r,
// but the range loop is still iterating over the original slice t
