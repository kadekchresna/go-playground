package main

import "fmt"

func main() {

	fmt.Println(freqcount([]int{1, 2, 2, 3, 1, 1}))

}

func freqcount(input []int) map[int]int {

	var res = map[int]int{}
	for _, v := range input {
		res[v]++
	}

	return res
}
