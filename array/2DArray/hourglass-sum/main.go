package main

import "fmt"

// There are 16 hourglasses in a 6x6 array. The  is the sum of the values in an hourglass.
// Calculate the hourglass sum for every hourglass in arr, then print the maximum hourglass sum.

// ex:
// -9 -9 -9  1 1 1
//  0 -9  0  4 3 2
// -9 -9 -9  1 2 3
//  0  0  8  6 6 0
//  0  0  0 -2 0 0
//  0  0  1  2 4 0

// The 16 hourglass sums are:

// -63, -34, -9, 12,
// -10,   0, 28, 23,
// -27, -11, -2, 10,
//   9,  17, 25, 18

/**
 * Geetting sum of hour glass
 * 1 1 1
 *   1		= 7
 * 1 1 1
 */

func main() {
	fmt.Println(hourglassSum([][]int{{0, -4, -6, 0, -7, -6}, {-1, -2, -6, -8, -3, -1}, {-8, -4, -2, -8, -8, -6}, {-3, -1, -2, -5, -7, -4}, {-3, -5, -3, -6, -6, -6}, {-3, -6, 0, -8, -6, -7}}))
}

func hourglassSum(arr [][]int) int {

	var maxTotal int = -99

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			currentTotal := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if maxTotal < currentTotal {
				maxTotal = currentTotal
			}
		}
	}
	return maxTotal
}
