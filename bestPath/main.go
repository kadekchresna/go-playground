package main

import "fmt"

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 0, 0, 1, 0}))
}

func jumpingOnClouds(c []int32) int32 {
	var path = []int32{}
	var iPath = []int{0}
	var l = 0

	for {
		for j := 2; j >= 1; j-- {
			if l+j >= len(c) {
				continue
			}
			if c[l+j] != 1 {
				path = append(path, c[l+j])
				iPath = append(iPath, l+j)
				l += j
				break
			}
		}
		fmt.Println("--", l)
		if l == len(c)-1 {
			break
		}
	}

	fmt.Println(path)
	fmt.Println(iPath)
	return int32(len(path))

}
