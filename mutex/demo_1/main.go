package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mt sync.Mutex
var v int

func main() {

	fmt.Println("Mutex")
	counter := 100
	wg.Add(counter)

	for i := 0; i < counter; i++ {
		go func() {
			fmt.Printf("before: %d\n", v)

			defer wg.Done()
			defer mt.Unlock()
			mt.Lock()
			v++
			fmt.Printf("after: %d\n", v)

		}()
	}
	wg.Wait()
	fmt.Println(v)

}
