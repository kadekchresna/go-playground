package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		fmt.Println("go")
		time.Sleep(3 * time.Second)
		ch <- "one"
		fmt.Println("done")

	}()

	select {
	case s := <-ch:
		fmt.Println()
		fmt.Println(s)
		fmt.Println()

	case sh := <-time.After(2 * time.Second):
		fmt.Println()
		fmt.Println("got timeout at", sh)
		fmt.Println()

	}

	fmt.Println("finish")
}
