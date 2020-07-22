package main

import (
	"fmt"
	"math"
)

func main() {
	var word string
	fmt.Println("Palindrome word check")
	fmt.Print("Input some word: ")
	fmt.Scanf("%s", &word)

	fmt.Println(len(word))

	if len(word) < 3 {
		fmt.Println("it must be 3 character or more.")
	}

	middle := math.Round(float64(len(word)) / 2.0)
	palindrome := true

	for i := 0; i < int(middle); i++ {
		if word[i] != word[(len(word)-1)-i] {
			palindrome = false
			break
		}
	}

	if palindrome {
		fmt.Printf("%s is a palindrome\n", word)
	} else {
		fmt.Printf("%s is not a palindrome\n", word)
	}

}
