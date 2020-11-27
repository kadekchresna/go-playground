package main

import (
	"fmt"
)

func main() {
	// fmt.Println(unique([]string{"a", "aa", "b", "c", "a"}))
	fmt.Println(uniqueArray("strings"))
}

func uniqueArray(str string) []string {
	var tmpM = make(map[rune]bool)
	var list = make([]string, 0)
	for _, val := range str {
		if _, value := tmpM[val]; !value {
			tmpM[val] = true
			list = append(list, string(val))
		}
	}
	return list
}
