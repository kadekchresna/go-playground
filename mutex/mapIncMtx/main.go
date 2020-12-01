package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
 * using sync.RWMutex for performace pusposes, sync.RWMutex will allow other go routines to write wile otherr read it.
 * im using channel for blocking the main routine to finish before other go routine complete its tasks
 */

type MapCounter struct {
	i map[string]int
	sync.RWMutex
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)
	mc := MapCounter{i: make(map[string]int)}

	go mc.Increment(10)
	go mc.ReadMap(ch, 10)

	fmt.Println(<-ch)

	// go func(c <-chan string) {
	// 	fmt.Println(<-c)
	// 	close(ch)
	// }(ch)
}

func (c *MapCounter) Increment(n int) {

	for i := 0; i < n; i++ {
		c.Lock()
		c.i[fmt.Sprintf("%d", i)] = i * 10
		c.Unlock()
	}

}

func (c *MapCounter) ReadMap(ch chan<- string, n int) {

	var i int
Loop:
	for {
		c.RLock()
		v := rand.Intn(n)
		if val, ok := c.i[fmt.Sprintf("%d", v)]; ok {
			fmt.Println(val)
			delete(c.i, fmt.Sprintf("%d", v))
			i++
		}

		c.RUnlock()
		if i == n {
			break Loop
		}
	}

	ch <- "done"
}
