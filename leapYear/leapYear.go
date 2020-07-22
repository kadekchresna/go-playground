package main

import "fmt"

func main() {
	var year int
	fmt.Println("Eyyy")
	fmt.Print("Enter any year: ")
	fmt.Scanf("%d", &year)

	leap := ((year%4 == 0 && year%100 != 0) || year%400 == 0)
	if leap {
		fmt.Println("Leap Year it is!")
	} else {
		fmt.Println("No, No i don't think i will.")
	}
}
