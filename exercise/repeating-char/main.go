package main

import "fmt"

func main() {

	fmt.Println(repeatingChar("abcabcbb"))

}

func repeatingChar(s string) int {
	var total int
	check := map[string]bool{}
	for _, v := range s {
		if _, ok := check[string(v)]; ok {
			return total
		}

		check[string(v)] = true

		total++
	}

	return total
}
