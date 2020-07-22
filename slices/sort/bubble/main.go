package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bubble Sort")
	s := []int{
		18, 20, 14, 17, 13,
		16, 11, 15, 12, 19, 13,
	}
	l := 0
	r := 1
	order := 1
	breakFlag := true
	breakCount := 0
	i := 0
	for {
		breakFlag = true
		if r == len(s) {
			l = 0
			r = 1
		}
		if breakCount == len(s) {
			break
		}
		fmt.Println(i, s)
		if ((s[r] - s[l]) * order) > 0 {
			s[l], s[r] = s[r], s[l]
			breakFlag = false
			breakCount = 0
		}

		l++
		r++
		if breakFlag {
			breakCount++
		}
		i++

	}

	fmt.Println("Sorted")
	fmt.Println(s)
}
