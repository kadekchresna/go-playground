package main

import "fmt"

func main() {

	addOne := func() func() int {
		i := 0

		return func() int {
			i++
			fmt.Println(i)
			return i
		}
	}
	o := addOne()
	fmt.Printf("%x\t", &addOne)
	fmt.Printf("%x", &o)
	fmt.Println()
	fmt.Println(o())
	fmt.Println(o())
	fmt.Println(o())
	fmt.Println(o())
}
