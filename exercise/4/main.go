package main

import "fmt"

func main() {
	fmt.Println(seatsBlockedBehind(16, 11, 5, 3))
}

func seatsBlockedBehind(nCols int, nRows int, col int, row int) int {
	return (nRows - row) * (nCols - (col - 1))
}
