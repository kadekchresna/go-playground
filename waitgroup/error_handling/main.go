package main

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

func main() {
	var wg sync.WaitGroup
	out := make(chan int)
	out2 := make(chan error)

	// ch := make(chan int)
	fmt.Println("BOO")
	// go foo(ch)
	wg.Add(2)

	go func() {
		// func foo(c chan int) {
		// c <- 1
		fmt.Println("FOO")
		out2 <- nil
		wg.Done()
	}()
	go func() {
		err := returnError()
		if err != nil {
			out2 <- err
		}
		fmt.Println("FOO")
		wg.Done()
	}()
	go func() {
		fmt.Println("POO")

	}()
	wg.Wait()
	close(out)
	// <-ch
	select {
	case err := <-out2:
		close(out2)
		if err != nil {
			panic(fmt.Sprintf("Error: %s", err))

		}
	}
}

func returnError() error {
	return errors.New("Error example")
}
