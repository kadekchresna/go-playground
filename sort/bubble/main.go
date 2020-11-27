package main

import "fmt"

func main() {
	bubble([]int32{3, 2, 1})
}

func bubble(a []int32) {
	var count int
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				count++
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Println(fmt.Sprintf("Array is sorted in %d swaps.", count))
	fmt.Println(fmt.Sprintf("First Element: %d", a[0]))
	fmt.Println(fmt.Sprintf("Last Element: %d", a[len(a)-1]))
}
