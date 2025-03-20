package main

import (
	"fmt"
	"reflect"
)

func checkType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
}

func main() {
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	checkType(arr)   // Output: Type: [3]int, Kind: array
	checkType(slice) // Output: Type: []int, Kind: slice
}
