package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	out := make(chan int)

	// ch := make(chan int)
	fmt.Println("BOO")
	// go foo(ch)
	wg.Add(1)

	go func() {
		// func foo(c chan int) {
		// c <- 1
		out <- 1
		fmt.Println("FOO")
		wg.Done()
	}()
	go func() {
		<-out
		fmt.Println("POO")
		wg.Wait()
		close(out)
	}()
	// <-ch
}
