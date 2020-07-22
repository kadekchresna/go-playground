package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minInput = 10
	maxInput = 12
)

func main() {
	fmt.Println("Hope you are lucky as you think you are!")
	fmt.Println("----------------------------------------")
	var inputedValue int

	for {
		fmt.Printf("Only number between %d and %d\n", minInput, maxInput)
		fmt.Print("Input some number: ")
		fmt.Scanf("%d", &inputedValue)
		fmt.Println()
		if inputedValue < minInput || inputedValue > maxInput {
			fmt.Printf("You cannot use %d", inputedValue)
		} else {
			break
		}
	}

	rand.Seed(time.Now().UnixNano())
	systemRandomVal := minInput + rand.Intn(maxInput-minInput)

	// fmt.Printf("yourRandomVal=%d systemRandomVal=%d",
	// 	yourRandomVal, systemRandomVal)

	yourRandomValSumOfDigits :=
		(inputedValue / 10) + (inputedValue % 10)
	systemRandomValSumOfDigits :=
		(systemRandomVal / 10) + (systemRandomVal % 10)

	// fmt.Printf("\nyourRandomVal=%d systemRandomVal=%d",
	// 	yourRandomValSumOfDigits, systemRandomValSumOfDigits)

	lotteryResult := ""
	if systemRandomVal == inputedValue {
		lotteryResult = "You won $1,000"
	} else if yourRandomValSumOfDigits == systemRandomValSumOfDigits {
		lotteryResult = "You won $500"
	} else {
		lotteryResult = "Try Again!"
	}

	fmt.Printf("\nyourRandomVal=%d systemRandomVal=%d \nlotteryResult=%s \n",
		inputedValue, systemRandomVal, lotteryResult)
}
