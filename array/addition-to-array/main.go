package main

import (
	"fmt"
)

// calculate array input with a number and transform the array matched with the result of the addition
// ex.
// input: [1, 2, 3, 4, 5], 3
// output: [1, 2, 3, 4, 8]

// input: [1, 2, 3, 4, 5], 11
// output: [1, 2, 3, 5, 6]

// input: [1, 2, 3, 4, 5], 11
// output: [1, 2, 3, 5, 6]

func main() {

	arraInput := []int{0, 2, 3, 4, 5}
	number := 100031

	fmt.Println(calculate(arraInput, number))
}

func power(input int, power int) int {

	result := 1
	for i := 0; i < power; i++ {
		result *= input
	}

	return result
}

func calculate(arraInput []int, number int) []int {

	if len(arraInput) == 0 {
		return arraInput
	}

	l := len(arraInput)

	result := arraInput[l-1] + number

	currentTotal := 0
	for i := 0; i < l-1; i++ {
		powerValue := l - i - 1
		currentTotal += arraInput[i] * power(10, powerValue)
	}

	currentTotal += result

	resArrRev := []int{}
	i := 0
	for {
		digitDiv := (currentTotal / power(10, i))
		if digitDiv < 10 {
			resArrRev = append(resArrRev, digitDiv)
			break
		}

		digit := digitDiv % 10
		i++

		resArrRev = append(resArrRev, digit)
	}

	resArr := []int{}
	for i := len(resArrRev) - 1; i >= 0; i-- {
		resArr = append(resArr, resArrRev[i])
	}

	return resArr
}
