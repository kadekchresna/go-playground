package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 7
)

func main() {
	fmt.Println("Try your luck.")
	fmt.Println("You got 5 lives. Use it wisely.")
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(max-min) + min

	lives := 5
	var input int
	for i := 0; i < lives; {

		fmt.Print("Input value: ")
		fmt.Scanf("%d\n", &input)
		if random > input {
			fmt.Println("Input bigger value.")
			i++
		} else if random < input {
			fmt.Println("Input smaller value.")
			i++
		} else {
			fmt.Printf("Value matched. it's %d\n", random)
			break
		}

	}
	fmt.Println("You died.")
}
