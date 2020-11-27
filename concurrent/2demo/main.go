package main

import (
	"fmt"
	"sync"
	"time"
)

type A struct {
	id int
}

func main() {
	start := time.Now()

	channel := make(chan A)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for a := range channel {
			wg.Add(1)
			go func(a A) {
				defer wg.Done()
				process(a)
			}(a)
		}

	}()

	for i := 0; i < 100; i++ {
		channel <- A{id: i}
	}
	close(channel)

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}

func process(a A) {
	fmt.Printf("Start processing %v\n", a)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Finish processing %v\n", a)
}
