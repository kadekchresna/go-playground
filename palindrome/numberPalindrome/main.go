package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(check(1100))

}

func check(num int) int {

	if num < 11 {
		return 11
	}
	s := fmt.Sprintf("%d", num)
	if Palindrome(s) {
		return num
	} else {
		iF, _ := strconv.Atoi(string(s[0]))
		iL, _ := strconv.Atoi(string(s[len(s)-1]))
		if iF < iL {
			return Inc(1, iF, s)
		} else {
			return Inc(0, iF, s)

		}
	}

}

func Inc(inc, iF int, s string) int {
	var result string
	for i := 0; i < len(s); i++ {
		result += fmt.Sprintf("%d", iF+inc)
	}
	res, _ := strconv.Atoi(result)
	return res

}

// func check(i int) int {
// 	num := i
// 	for {
// 		b := Palindrome(fmt.Sprintf("%d", num))
// 		if !b {
// 			num++
// 		} else {
// 			return num
// 		}
// 	}
// }

func Palindrome(s string) bool {
	var mid = len(s) / 2
	var l = 0
	var r = len(s) - 1
	var result = true

	for {
		if s[l] != s[r] {
			result = false
		}
		l++
		r--
		if l > mid {
			return result
		}
	}
}
