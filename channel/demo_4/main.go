package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(1)
	go func() {
		// out <- 2
		wg.Done()
	}()

	wg.Wait()
	// close(out)

	fmt.Println(<-out)
}
