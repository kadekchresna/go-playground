package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(1 * time.Second)
	go counter(t)
	time.Sleep(10 * time.Second)
}

func counter(t *time.Ticker) {

	i := 0

	for t := range t.C {
		time.Sleep(10 * time.Second)
		i++
		fmt.Println("Count ", i, " at ", t)
	}
}
