package main

import "fmt"

func main() {

	PrintSlice([]int{1, 2, 3, 4})
	PrintSlice[int]([]int{1, 2, 3, 4})
	PrintSlice[float64]([]float64{1.4, 2.52, 0.3, 1.4})

}

func PrintSliceAny[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func PrintSlice[T int | float64](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
