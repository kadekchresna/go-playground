package main

import "fmt"

// detect palindrome of an integer but without using any convertion
// 121 -> true
// 1221 -> true
// 123 -> false
// 122 -> false
// 1 -> true
// 11 -> true
// 12 -> false

func main() {

	fmt.Println(CheckPalindrome(11))

}

func power(n int, power int) int {
	result := 1

	for i := 0; i < power; i++ {
		result *= n
	}

	return result
}

func CheckPalindrome(n int) bool {

	isPalindrom := false

	valueInSlice := []int{}
	for i := 0; ; i++ {
		pow := power(10, i)
		devRemainder := n / pow

		modReminder := devRemainder % 10

		if (n - pow) < 0 {
			break
		}
		valueInSlice = append(valueInSlice, modReminder)

	}

	lenResult := len(valueInSlice)
	if lenResult < 2 {
		return isPalindrom
	}

	halfOfValue := lenResult / 2
	for i := 0; i < halfOfValue; i++ {
		if valueInSlice[i] != valueInSlice[lenResult-(i+1)] {
			isPalindrom = false
			break
		}

		isPalindrom = true
	}

	return isPalindrom
}
