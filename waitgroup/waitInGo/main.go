package main

import (
	"fmt"
	"sync"
)

type waiter struct {
	wg sync.WaitGroup
}

func main() {

	fmt.Println("a")

	wait := waiter{
		wg: sync.WaitGroup{},
	}
	wait.wg.Add(1)
	go wait.Foo()
	wait.wg.Wait()
	fmt.Println("c")
}

func (c *waiter) Foo() {
	c.wg.Add(1)
	go func() {
		fmt.Println("foo")
		c.wg.Done()
	}()

	go func() {
		fmt.Println("waited")
		c.wg.Wait()
	}()

	c.wg.Done()

}
