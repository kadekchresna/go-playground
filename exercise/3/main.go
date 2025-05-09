package main

import "fmt"

func main() {
	fmt.Println(solution(3, 10))
}

func solution(n int, m int) int {
	return m - (m % n)
}
