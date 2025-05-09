package main

import "fmt"

func main() {
	// fmt.Println(minimumSwaps([]int32{2, 3, 4, 1, 5}))
	fmt.Println(minimumSwaps([]int32{1, 3, 5, 2, 4, 6, 7}))
	// fmt.Println(minimumSwaps([]int32{4, 3, 1, 2}))
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swap int32

	var i int32 = 0

	for i = 0; i < int32(len(arr)); i++ {
		if arr[i] != i+1 {
			diff := arr[i] - 1
			temp := arr[diff]

			arr[diff] = arr[i]
			arr[i] = temp
			swap++
			i--
		}
	}

	return swap
}
