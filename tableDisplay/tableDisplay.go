package main

import "fmt"

func main() {
	fmt.Println("i\t\tt*2\t\tt*3")
	fmt.Println("==\t\t===\t\t===")
	tables(4)
}

func tables(u int) {
	for i := 0; i <= u; i++ {
		fmt.Printf("%d\t\t%d\t\t%d\n", i, i*i, i*i*i)
	}
}
