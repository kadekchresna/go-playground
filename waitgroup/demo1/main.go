package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("a")
	fmt.Println(foo())
	fmt.Println("c")
}

func foo() <-chan int {
	var wg sync.WaitGroup
	wg.Add(1)
	out := make(chan int)

	go func() {
		fmt.Println("foo")
		wg.Done()

	}()
	wg.Wait()

	go func() {
		wg.Wait()
		close(out)

	}()

	return out

}
