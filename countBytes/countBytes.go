package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Hello"
	s2 := "𠜎ち"

	fmt.Println([]int32(s1))
	fmt.Println([]int32(s2))

	fmt.Println([]byte(s1))
	fmt.Println([]byte(s2))

	fmt.Printf("%s, %d, %d\n", s1, len(s1), utf8.RuneCountInString(s1))
	fmt.Printf("%s, %d, %d\n", s2, len(s2), utf8.RuneCountInString(s2))

	var sliceRune []rune

	for _, val := range s1 {
		sliceRune = append(sliceRune, val)
	}

	fmt.Printf("%v\n", sliceRune)
	fmt.Printf("%q\n", sliceRune)
	fmt.Printf("%q\n", string(sliceRune))
}
