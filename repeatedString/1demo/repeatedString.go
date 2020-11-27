package main

import (
	"fmt"
	"strings"
)

/**
 * Count how many "a" inside repeated string s n times
 */

func main() {
	fmt.Println(repeatedString("aba", 10))
}

func repeatedString(s string, n int64) int64 {

	left := n % int64(len(s))

	count := int64(strings.Count(s, "a")) * (n / int64(len(s)))
	return count + int64(strings.Count(s[0:left], "a"))
}
