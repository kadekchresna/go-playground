package main

import "fmt"

func main() {
	// Highest([]int{1, 2, 3, 8, 9, 3, 2, 1})
	fmt.Println(Highest([]int{1, 2, 1, 2, 2, 1}))
	// Highest([]int{7, 1, 2, 9, 7, 2, 1})
}

func Highest(arr []int) int {
	var highest int
	var middle = len(arr) / 2
	var arrLeft = arr[:middle]
	var arrRight = arr[middle:]
	var l = 0
	var r = len(arrRight) - 1
	for {
		if arrLeft[l] == arrRight[r] && arrLeft[l+1] == arrRight[r-1] {
			highest = arrLeft[l+1]
			l++
		}

		r--
		if r == 1 {
			r = len(arrRight) - 1
			l++
		}

		if l > len(arrLeft)-2 {
			break
		}

	}
	return highest
}
