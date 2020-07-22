package main

import "fmt"

func main() {
	var height float32
	var weight float32
	fmt.Print("Input Weight: ")
	fmt.Scanf("%v", &weight)
	fmt.Print("\nInput Height: ")
	fmt.Scanf("%v", &height)

	bmi := weight / (height * height)

	switch {
	case bmi < 18.5:
		fmt.Println("Undeweight")
	case bmi < 25:
		fmt.Println("Normal")
	case bmi < 30:
		fmt.Println("Overweight")
	default:
		fmt.Println("Obese")

	}

}
