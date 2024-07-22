// File name: ...\s04\exercises\set-3\bubble_sort
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

/*
Exercise: Write a program to sort a list of integers
using 'Bubble Sort'.
Also add a flag to your program to switch between
ascending and descending sorts.

Bubble sort:
- probably the simplest sorting algorithm;
- bubble up the largest (or the smallest); then the second largest, ...
- performance: O(n2) [for best-case, average-case and worst-case]

Steps: 4,2,1,3 --> 2,4,1,3 -->
	   2,1,4,3 --> 2,1,3,4 --> 1,2,3,4

*/

func main() {
	items := []int{
		18, 20, 14, 17, 13,
		16, 11, 15, 12, 19,
	}

	order := 1 // 1 for ascending and -1 for descensing

	end := len(items)
	ii := 0

	for i := 0; i < end; i++ {
		for j := 0; j < (end - 1 - i); j++ {
			fmt.Println(ii)
			// if items[j] > items[j+1] {
			if ((items[j] - items[j+1]) * order) > 1 {
				items[j], items[j+1] = items[j+1], items[j]
			}
			ii++
		}
	}

	fmt.Printf("%v\n", items)
}
