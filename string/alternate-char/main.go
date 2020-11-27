package main

import "fmt"

/**
 * Alternating Characters
 * https://s3.amazonaws.com/hr-assets/0/1502450721-a0a2e9b5bd-alternatingCharacter2.png
 * reeturn how much deletion
 */

func main() {
	fmt.Println(countDel("AAAA"))
	fmt.Println(countDel("BBBBB"))
	fmt.Println(countDel("ABABABAB"))
	fmt.Println(countDel("AAABBB"))
}

func countDel(s string) int32 {
	var result int32
	var i int
	var l = 1
	for {
		if s[i] == s[l] {
			l++
			result++
		} else {
			i = l
			l = i + 1
		}

		if l >= len(s) {
			break
		}
	}
	return result
}
