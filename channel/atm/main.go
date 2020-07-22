package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var currentBalance = 200

var wg sync.WaitGroup

func main() {
	fmt.Println(<-doATM())
}

func doATM() chan int {
	c := make(chan int)

	// credit := make(chan int)
	// debit := make(chan int)

	wg.Add(3)
	go credit(c)
	go debit(c)
	go balance(c)
	wg.Wait()
	close(c)
	return c
}
func credit(c chan int) {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("credit")
	b := rand.Intn(10)
	oldB := currentBalance
	currentBalance = currentBalance + b
	fmt.Println(currentBalance)
	fmt.Printf("%v + %v = %v\n", oldB, b, currentBalance)

	c <- currentBalance
	wg.Done()

}
func debit(c chan<- int) {
	rand.Seed(time.Now().UnixNano())

	b := rand.Intn(10)
	oldB := currentBalance
	currentBalance = currentBalance + (-1 * b)
	fmt.Println("debit")

	fmt.Printf("%v - %v = %v\n", oldB, b, currentBalance)
	c <- currentBalance
	wg.Done()

}
func balance(c <-chan int) {
	fmt.Println("balance")

	num, ok := <-c
	if !ok {
		fmt.Println("Error:")
	}
	fmt.Println(num)
	// for i := 0; i < 2; i++ {
	// 	num, ok := <-c
	// 	if !ok {
	// 		fmt.Println("Error:")
	// 	}
	// 	fmt.Println(num)
	// }
	wg.Done()

}
