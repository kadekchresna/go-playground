package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 *
 * select channel block, so when we check channel we used select. when we use it we need to add default if there are next flow after select
 */

func main() {

	ch := make(chan string)
	f := make(chan string)
	go check(ch, f)
	// go set(ch)
	// <-f
	time.Sleep(10 * time.Second)
}

func set(c chan string) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		c <- "stringgggg"
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()
}

func check(c chan string, f chan string) {
	select {
	case <-c:
		// fmt.Println(<-c)
		f <- "done"
		// default:
		// fmt.Println("thers notin")
		// f <- "done"

	}

	fmt.Println("donee cheeck")
}
