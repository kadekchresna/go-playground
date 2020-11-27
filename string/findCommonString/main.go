package main

import "fmt"

/**
 * Given two strings, determine if they share a common substring. A substring may be as small as one character.
 * For example, the words "a", "and", "art" share the common substring . The words "be" and "cat" do not share a substring.
 * It should return a string, either YES or NO based on whether the strings share a common substring.
 */

func main() {
	fmt.Println(twoStrings("hello", "world"))
	fmt.Println(twoStrings("hi", "world"))
}

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	var mapS = make(map[string]bool)
	var common bool
	for _, sub := range s1 {
		if _, ok := mapS[string(sub)]; !ok {
			mapS[string(sub)] = true
		}
	}

	for _, sub := range s2 {
		if _, ok := mapS[string(sub)]; ok {
			common = true
		}
	}

	if common {
		return fmt.Sprintf("YES")
	} else {
		return fmt.Sprintf("NO")
	}

}
