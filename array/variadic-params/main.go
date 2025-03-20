package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{3, 4}
	c := append(a, b...) // Merging slices
	fmt.Println(c)
}
