package main

import "fmt"

func main() {
	fmt.Println(solution(10, 2, 5))
	// fmt.Println(solution(12, 5, 9))
}

func solution(n int, firstNumber int, moveRight int) int {
	i := firstNumber
	loop := 0
	for {
		loop++
		if i == n-1 {
			i = 0
			continue
		}
		i++
		if loop >= moveRight {
			break
		}

	}

	return i
}
