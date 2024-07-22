package main

import "fmt"

/**
 * Two strings are anagrams of each other if the letters of one string can be rearranged to form the other string.
 * Given a string, find the number of pairs of substrings of the string that are anagrams of each other.
 *
 * For example s=mom, the list of all anagrammatic pairs is [m,m], [mo,om] at positions [[0],[2]], [[0,1], [1,2]]  respectively.
 */

func main() {
	fmt.Println(sherlockAndAnagrams("mom"))
}

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	// var anagrams = make([]string, 0)
	// var anaMap = make(map[string]int)
	// var anaIMap = make(map[string]int)
	var multi bool
	// var i int
	// for _, sub := range s {
	// 	if _, ok := anaMap[string(sub)]; ok {
	// 		if val, ok := anaMap[string(sub)]; ok {

	// 		}
	// 		multi = true
	// 		anaIMap[string(sub)] = i
	// 	} else {
	// 		anaMap[string(sub)] = i
	// 	}
	// 	i++
	// }

	// for _, val := range anaIMap {

	// }
	if multi {
		return 3
	} else {
		return 0
	}
}
