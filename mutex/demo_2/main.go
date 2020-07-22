package main

import (
	"fmt"
	"sync"
)

type safeNumber struct {
	mt  sync.Mutex
	val int
}

func (s *safeNumber) Get() int {
	defer s.mt.Unlock()
	s.mt.Lock()
	return s.val
}
func (s *safeNumber) Set(v int) {
	defer s.mt.Unlock()
	s.mt.Lock()
	s.val = v
}
func (s *safeNumber) Inc() {
	defer s.mt.Unlock()
	s.mt.Lock()
	s.val++
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


