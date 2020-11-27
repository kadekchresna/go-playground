package main

import "fmt"

/**
 * Geetting sum of hour glass
 * 1 1 1
 *   1		= 7
 * 1 1 1
 */

func main() {
	fmt.Println(hourglassSum([][]int32{{0, -4, -6, 0, -7, -6}, {-1, -2, -6, -8, -3, -1}, {-8, -4, -2, -8, -8, -6}, {-3, -1, -2, -5, -7, -4}, {-3, -5, -3, -6, -6, -6}, {-3, -6, 0, -8, -6, -7}}))
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var hGlass int32
	hGlass = -100000000
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			total := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if hGlass <= total {
				hGlass = total
			}
		}
	}
	return hGlass
}

// 1 1 1 0 0 0
// 0 1 0 0 0 0
// 1 1 1 0 0 0
// 0 0 2 4 4 0
// 0 0 0 2 0 0
// 0 0 1 2 4 0

// [0][0]
// [0][1]
// [0][2]
// [1][1]
// [2][0]
// [2][1]
// [2][2]

// [0][1]
// [0][2]
// [0][3]
// [1][2]
// [2][1]
// [2][2]
// [2][3]
