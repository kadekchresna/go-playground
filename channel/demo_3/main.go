package main

import (
	"fmt"
	"sync"
)

func main() {

	out := foo()
	fmt.Println(<-out)

}

func foo() <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// go func() {
	// 	fmt.Println("here is wait ")
	// 	wg.Wait()
	// 	close(out)
	// }()
	wg.Add(1)
	go func() {
		fmt.Println("here is foo ")
		// out <- 2
		wg.Done()
	}()

	wg.Wait()
	close(out)

	return out

}
