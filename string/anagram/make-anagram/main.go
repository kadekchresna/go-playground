package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeAnagram("showman", "woman"))
	fmt.Println(makeAnagram("woman", "shoooowman"))
	fmt.Println(makeAnagram("cde", "abc"))
	// fmt.Println(makeAnagram("fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke"))
}
func makeAnagram(a string, b string) int32 {
	var aMap = make(map[string]int)
	var tmpMap = make(map[string]int)
	var result int32
	for _, s := range a {
		if aMap[string(s)] == 0 {
			aMap[string(s)] = 1
			tmpMap[string(s)] = 1
		} else {
			aMap[string(s)] += 1
			tmpMap[string(s)] += 1
		}
	}

	// fmt.Println(aMap)
	for _, s := range b {
		if value, ok := aMap[string(s)]; ok {
			if value > 0 {
				aMap[string(s)] -= 1
			} else {
				result++
			}
		} else {
			result++
		}

	}
	// result = len(b) - result
	for key, value := range aMap {
		if tmpMap[key] != value && value != 0 {
			result += int32(value)
		}
		if tmpMap[key] == value {
			result += int32(value)
		}
	}
	// fmt.Println(aMap)
	// fmt.Println(tmpMap)

	return result
}
