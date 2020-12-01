package main

import (
	"fmt"
	"sync"
)

/**
 * increment with locking variable might get race between increment and decrement we use mutex,
 * when getting the value result inorder to avoid the race condition we can use sync flow with wait group
 *
 */

type safeCounter struct {
	i int
	sync.Mutex
	sync.WaitGroup
}

func main() {
	sc := new(safeCounter)

	for i := 0; i < 100; i++ {
		sc.Add(2)
		go sc.Increment()
		go sc.Decrement()
	}

	sc.Wait()

	fmt.Println(sc.GetValue())

}

func (c *safeCounter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.i++

	c.Done()
}

func (c *safeCounter) Decrement() {
	c.Lock()
	defer c.Unlock()
	c.i--

	c.Done()
}

func (c *safeCounter) GetValue() int {
	return c.i
}
