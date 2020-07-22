package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := gen(2, 5)

	// fmt.Println(<-out)

	// for n := range sq(sq(ch)) {
	// 	fmt.Println(n)
	// }

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(ch)
	c2 := sq(ch)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}

}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n

		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	// output := func(c <-chan int) {
	// 	for n := range c {
	// 		out <- n
	// 	}
	// 	wg.Done()
	wg.Add(len(cs))
	// }
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()

		}(c)
	}
	// wg.Wait()

	// close(out)
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
