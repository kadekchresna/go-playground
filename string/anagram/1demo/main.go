package main

import "fmt"

/**
 * Funsi check data anagram dari 2 kata
 */

func main() {
	fmt.Println(compare("yoo", "oyo"))
}

func compareV2(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var result = true
	var sOMap = make(map[string]int)

	for _, charr := range s1 {
		if _, ok := sOMap[string(charr)]; ok {
			sOMap[string(charr)] += 1
		} else {
			sOMap[string(charr)] = 1
		}
	}

	for _, charr := range s2 {
		if _, ok := sOMap[string(charr)]; ok {
			sOMap[string(charr)] -= 1
		}
	}

	for _, val := range sOMap {
		if val > 0 {
			result = false
		}
	}
	return result
}

func compare(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var result = true
	var sOMap = make(map[string]int)

	for _, charr := range s1 {
		if _, ok := sOMap[string(charr)]; ok {
			sOMap[string(charr)] += 1
		} else {
			sOMap[string(charr)] = 1
		}
	}

	for _, charr := range s2 {
		if _, ok := sOMap[string(charr)]; ok {
			sOMap[string(charr)] -= 1
		}
	}

	for _, val := range sOMap {
		if val > 0 {
			result = false
		}
	}
	return result
}
