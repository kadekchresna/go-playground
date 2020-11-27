package main

import "fmt"

/**
 * You are given an unordered array consisting of consecutive integers  E[1, 2, 3, ..., n] without any duplicates.
 * You are allowed to swap any two elements.
 * You need to find the minimum number of swaps required to sort the array in ascending order.
 */

func main() {
	fmt.Println(minimumSwaps([]int32{2, 3, 4, 1, 5}))
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var iSorted = 0
	var smallest = 0
	var i = 1
	var swapped int32
	for iSorted < len(arr) {
		if arr[smallest] > arr[i] {
			smallest = i
		}
		i++

		if i > len(arr)-1 {
			if arr[iSorted] != arr[smallest] {
				swap(&arr[iSorted], &arr[smallest])
				swapped++
			}
			iSorted++
			i = iSorted
			smallest = i
		}
	}
	return swapped

}

func swap(a *int32, b *int32) {
	*a, *b = *b, *a
}
