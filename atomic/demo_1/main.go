package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type safeNumber struct {
	val int64
}

func (s *safeNumber) Get() int64 {
	return atomic.LoadInt64(&s.val)
}
func (s *safeNumber) Set(v int64) {
	atomic.SwapInt64(&s.val, v)
}
func (s *safeNumber) Inc() {
	atomic.AddInt64(&s.val, 1)
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	sn := safeNumber{}
	go func() {
		defer wg.Done()
		sn.Set(2)
	}()
	wg.Wait()

	wg.Add(1)

	go func() {
		defer wg.Done()
		sn.Set(4)
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sn.Set(5)
	}()
	wg.Wait()

	wg.Add(1)

	go func() {
		defer wg.Done()
		sn.Inc()
	}()
	wg.Wait()
	fmt.Println(sn.Get())
}
