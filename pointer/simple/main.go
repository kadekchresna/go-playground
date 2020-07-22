package main

import "fmt"

func main() {
	x := []int{2, 1, 3, 4, 10, 200}
	p := &x
	fmt.Printf("%x, %d, %b", p, p, p)
	fmt.Println()
}
