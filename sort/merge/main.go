package sorts

import (
	"fmt"
)

var swapped int

func main() {
	fmt.Println(MergeSort([]int{2, 3, 4, 1, 5}))
	fmt.Println(swapped)
}

func helper(arrLeft []int, arrRight []int) []int {
	var result []int
	var i = 0
	var j = 0

	for i < len(arrLeft) && j < len(arrRight) {
		swapped++
		if arrLeft[i] > arrRight[j] {
			result = append(result, arrRight[j])
			j++
		} else {
			result = append(result, arrLeft[i])
			i++
		}
	}

	if i < len(arrLeft) {
		result = append(result, arrLeft[i:]...)
	}
	if j < len(arrRight) {
		result = append(result, arrRight[j:]...)
	}
	return result
}

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:])
	return helper(left, right)

}
