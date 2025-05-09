package main

import (
	"fmt"
)

// Starting with a 1-indexed array of zeros and a list of operations,
// for each operation add a value to each array element between two given indices, inclusive.
// Once all operations have been performed, return the maximum value in the array.

// ex
//
// n=10
// queries =
//   a, b, k
// [[1,5,3], > first op
// [4,8,7],  > second op
// [6,9,1]]  > third op

// in range a - b, add k value to array of zero in size of n

// result
// [0,0,0,0,0,0,0,0,0,0] > initial
// [3,3,3,3,3,0,0,0,0,0] > after first op
// [3,3,3,10,10,7,7,7,0,0] > after second op
// [3,3,3,10,10,8,8,8,1,0] > after third op

// the highest value is 10

func main() {

	// fmt.Println(arrayManipulation(5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}))
	// fmt.Println(arrayManipulation(10, [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}}))
	fmt.Println(arrayManipulation(10, [][]int32{{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15}}))
	// [1,9,15,15,15,8,0,0,0,0]
	// fmt.Println(arrayManipulation(5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}))
}

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

// func arrayManipulation(n int32, queries [][]int32) int64 {
// 	// Write your code here

// 	var highest int64

// 	arrayRes := make([]int64, n)

// 	for _, q := range queries {

// 		bottom := q[0] - 1
// 		top := q[1] - 1
// 		value := q[2]

// 		for i := bottom; i <= top; i++ {
// 			arrayRes[i] += int64(value)

// 			if highest < arrayRes[i] {
// 				highest = arrayRes[i]
// 			}

// 		}

// 		fmt.Println(arrayRes)
// 	}

// 	return highest

// }

// func arrayManipulation(n int32, queries [][]int32) int64 {
// 	// Write your code here

// 	var highest int64
// 	mapIntersection := make(map[string]int64, len(queries))

// 	for _, q := range queries {

// 		newBottom := q[0] - 1
// 		newTop := q[1] - 1

// 		value := q[2]

// 		if len(mapIntersection) == 0 {

// 			mapIntersection[fmt.Sprintf("%d-%d", newBottom, newTop)] = int64(value)

// 			highest = int64(value)
// 		}

// 		for k, v := range mapIntersection {

// 			startEnd := strings.Split(k, "-")
// 			bottom, _ := strconv.ParseInt(startEnd[0], 10, 32)
// 			top, _ := strconv.ParseInt(startEnd[1], 10, 32)
// 			// handle in-range && inter-section
// 			if int32(bottom) <= newBottom && int32(top) >= newTop ||
// 				int32(bottom) >= newBottom && int32(bottom) <= newTop ||
// 				int32(bottom) >= newBottom && int32(bottom) >= newTop ||
// 				int32(bottom) <= newBottom && int32(bottom) <= newTop {

// 				fmt.Println("intersection")
// 				fmt.Println("bottom-top", bottom, top, int32(bottom), int32(top))
// 				fmt.Println("newBottom-newTop", newBottom, newTop)

// 				v += int64(value)
// 				bottom := newBottom
// 				top := newTop

// 				delete(mapIntersection, k)

// 				mapIntersection[fmt.Sprintf("%d-%d", bottom, top)] = v
// 				// handle outside
// 			} else {

// 				fmt.Println("outside")

// 				mapIntersection[fmt.Sprintf("%d-%d", bottom, top)] = v

// 			}

// 			if highest < v {
// 				highest = v
// 			}

// 		}

// 	}

// 	return highest

// }

func arrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here

	var highest int64
	var prefixSum int64

	diffArray := make([]int64, n+1)

	// do the difference array
	for _, q := range queries {

		bottom := q[0] - 1
		top := q[1] - 1
		value := q[2]

		diffArray[bottom] += int64(value)
		diffArray[top+1] -= int64(value)
	}

	// do the prefix sum to reconstruct the array
	for _, v := range diffArray {

		prefixSum += v
		if highest < prefixSum {
			highest = prefixSum
		}
	}

	return highest

}
