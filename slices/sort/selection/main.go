package main

import (
	"fmt"
	"math"
)

func main() {

	s := []int{14, 23, 5, 26, 78, 43, 73, 115, 263, 4, 20}

	selection(&s)
	fmt.Println(s)

	result := binarySearch(s, 23)
	fmt.Println(result)

}

func binarySearch(slice []int, value int) int {
	s := slice
	min := 0
	max := len(s) - 1
	match := false
	middle := 0
	i := 0
	for {
		middle = int(math.Round(float64((max + min) / 2)))

		if middle >= len(s) {
			break
		}
		max = len(s) - 1
		if value > s[middle] {
			s = s[middle:]
		} else if value < s[middle] {
			s = s[:middle+1]
		} else if value == s[middle] {
			match = true
			break
		}
		i++

	}
	fmt.Println(i)

	if match {
		fmt.Println("Match")
		return value
	}

	fmt.Println("Try different number")
	return 0
}

func selection(slice *[]int) {
	s := *slice
	// s := []int{
	// 	18, 20, 14, 17, 13,
	// 	16, 11, 15, 12, 19}
	// s := []int{14, 23, 5, 26, 78, 43, 73, 115, 263, 4, 20}
	// fmt.Println(s)
	// fmt.Println()
	lowest := s[0]
	sortedIndex := 0
	i := sortedIndex + 1
	count := 0
	for {
		if i == len(s) {
			s[sortedIndex] = lowest
			sortedIndex++
			if sortedIndex == len(s)-1 {
				break
			}
			lowest = s[sortedIndex]
			i = sortedIndex + 1

		}
		if lowest > s[i] {
			lowest, s[i] = s[i], lowest
		}

		i++
		count++
	}

	fmt.Println("------------------------------------")

	fmt.Println("Selection Sort")
	// fmt.Println(count)
	fmt.Println(s)
	fmt.Println("------------------------------------")
}
