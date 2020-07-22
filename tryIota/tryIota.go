package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = 2
		c3 = iota
		c4 = 9
		c5 = iota
	)
	const d1 = iota

	const (
		off, false, no = iota, iota, iota
		yes, true, on
	)

	const (
		ten    = (iota + 1) * 10
		twenty = (iota + 1) * 10
		thirty = (iota + 1) * 10
	)

	fmt.Printf("thenth: %d,%d,%d\n", ten, twenty, thirty)

	fmt.Printf("%d,%d,%d,%d,%d\n", c1, c2, c3, c4, c5)
	fmt.Printf("%d\n", d1)

	fmt.Printf("falsy: %d,%d,%d\n", off, false, no)
	fmt.Printf("truthy: %d,%d,%d\n", yes, true, on)
}
