package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution("a"))
}

// func solution(year int) int {
// 	cen := math.Floor(float64(year / 100))
// 	if int(cen*100) < year {
// 		cen++
// 	}

// 	return int(cen)
// }

func solution(inputString string) bool {

	if len(inputString) < 2 {
		return false
	}

	if len(inputString)%2 == 0 {
		return false
	}
	if inputString[len(inputString)/2+1:len(inputString)] == inputString[0:len(inputString)/2] {
		return true
	}

	return false

}
