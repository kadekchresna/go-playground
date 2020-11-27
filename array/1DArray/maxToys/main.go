package main

import "fmt"

/**
 * get max toys with min budget
 */

func main() {
	fmt.Println(maximumToys([]int32{1, 12, 5, 111, 200, 1000, 10}, 50))
	fmt.Println(maximumToys([]int32{1, 2, 3, 4}, 7))
}
func maximumToys(prices []int32, k int32) int32 {

	var tmp int32
	var toys int32

	sorted := merge(prices)
	fmt.Println(sorted)

	for _, val := range sorted {
		if tmp+val <= k {
			tmp += val
			toys++
		}
	}
	return toys
}

func helper(arrL, arrR []int32) []int32 {
	var res []int32
	var l int32
	var r int32 = 0

	for l < int32(len(arrL)) && r < int32(len(arrR)) {
		if arrL[l] < arrR[r] {
			res = append(res, arrL[l])
			l++
		} else {
			res = append(res, arrR[r])
			r++
		}
	}

	if l < int32(len(arrL)) {
		res = append(res, arrL[l:]...)
	}
	if r < int32(len(arrR)) {
		res = append(res, arrR[r:]...)
	}
	return res
}

func merge(arr []int32) []int32 {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left := merge(arr[:mid])
	right := merge(arr[mid:])
	return helper(left, right)
}
