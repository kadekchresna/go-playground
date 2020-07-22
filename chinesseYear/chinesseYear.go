package main

import "fmt"

func main() {
	fmt.Println("aaaaaa")

	y := 2040

	switch y % 12 {
	case 0:
		fmt.Println("Ye")
	default:
		fmt.Println("Ya")

	}

}
