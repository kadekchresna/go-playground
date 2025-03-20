package main

import "fmt"

type Person struct {
	Name string
}

func main() {

	PrintSlice([]int{1, 2, 3, 4})
	PrintSlice[int]([]int{1, 2, 3, 4})
	PrintSlice[Person]([]Person{{"Name"}})
	PrintSlice[[3]int]([][3]int{{1, 2, 3}})
	PrintSlice[[3]Person]([][3]Person{{{Name: "name"}}})
	PrintSlice[float64]([]float64{1.4, 2.52, 0.3, 1.4})

	// ex with types that are not comparable

	// []int does not satisfy comparable
	// PrintSlice[[]int]([][]int{{1,2,3}})

	// map[string]int does not satisfy comparable
	// PrintSlice[map[string]int]([]map[string]int{{}})

	PrintKeyValue(map[string]float64{"1": 2.4})

}

func PrintSlice[T comparable](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func PrintKeyValue[K comparable, V int | float64](m map[K]V) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
