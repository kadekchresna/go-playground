package main

import "fmt"

func main() {

	for i := 0; i <= 4; i++ {
		fmt.Println(fibonacci(i))

	}
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
func fibonacci(l int) int {
	if l == 0 {
		return 0
	} else if l == 1 {
		return 1
	} else {
		return fibonacci(l-1) + fibonacci(l-2)
	}

}
