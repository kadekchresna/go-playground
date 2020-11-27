package main

import "fmt"

func main() {
	print(20)
}

func print(num int) {
	for i := 1; i < num; i++ {
		if i%3 == 0 || i == 1 {
			fmt.Printf("%d ", i)

			for j := 1; j < i; j++ {
				if (j+i)%2 == 0 {
					fmt.Printf("%d ", j+i)
				}
			}
			fmt.Println()

		}
	}
}
