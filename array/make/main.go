package main

import "fmt"

func main() {

	//             l  c
	// make([]int, 4, 4)
	// length (l) is the total data inside the map or slice
	// capacity (c) is the capacity of a slice
	// make only create slice not array
	// array can be created only through defining a variabel by [N]<type>

	arr := make([]int, 4, 4)

	fmt.Println(arr)

	s := make([]int, 0)
	// s := make([]int, 0, 3)
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
	}
}
