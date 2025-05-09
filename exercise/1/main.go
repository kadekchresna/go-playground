package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(solution(2024))

}

func solution(n int) int {
	s := strconv.Itoa(n)

	firstHalf, _ := strconv.Atoi(s[len(s)/2:])
	scdHalf, _ := strconv.Atoi(s[:len(s)/2])

	return firstHalf + scdHalf
}
