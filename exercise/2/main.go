package main

import (
	"fmt"
)

func main() {

	fmt.Println(solution(10))

}

func solution(n int) string {
	s := ""
	for i := 0; i < n; i++ {

		if i == n-1 {
			s += "1"
			break
		}
		s += "0"
	}

	return s
}
