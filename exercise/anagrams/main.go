package main

import "fmt"

func main() {

	fmt.Println(anagrams("listens", "silent"))

}

func anagrams(s string, t string) bool {

	var unique = map[string]int{}
	for _, v := range s {
		unique[string(v)]++
	}

	fmt.Println(unique)

	for _, v := range t {
		unique[string(v)]--
	}

	for _, v := range unique {
		if v > 0 {
			return false
		}
	}

	return true
}
