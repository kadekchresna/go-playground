package main

import "fmt"

/**
 * Print an integer denoting the minimum number of bribes needed to get the queue into its final state.
 * Print Too chaotic if the state is invalid, i.e. it requires a person to have bribed more than 2 people.
 */

func main() {
	minimumBribes([]int32{2, 1, 5, 3, 4})
}

// func minimumBribes(q []int32) {
// 	var bribes int
// 	for i := len(q) - 1; i >= 0; i-- {
// 		if q[i] != int32(i+1) {

// 			if (i-1) >= 0 && q[i-1] == int32(i+1) {
// 				bribes++
// 				q[i], q[i-1] = q[i-1], q[i]
// 			} else if (i-2) >= 0 && q[i-2] == int32(i+1) {
// 				bribes += 2
// 				q[i-2] = q[i-1]
// 				q[i-1] = q[i]
// 				q[i] = int32(i + 1)
// 			} else {
// 				fmt.Println("Too chaotic")
// 				return
// 			}
// 		}
// 	}
// 	fmt.Println(bribes)
// 	return
// }

func minimumBribes(q []int32) {
	// var bribes int32

	// for i := len(q); i > 0 ; i-- {

	// }

	for i := 0; ; i-- {

		fmt.Println(i)
		if i < -20 {
			break
		}
	}
}
