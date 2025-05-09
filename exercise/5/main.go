package main

import "fmt"

func main() {
	// fmt.Println(solution(3, 10))
	fmt.Println(solution(2, 20))
}

func solution(divisor int, bound int) int {
	res := bound % divisor

	if res > 0 {
		return bound - res
	}

	return bound
}
