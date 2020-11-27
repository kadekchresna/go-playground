package main

import (
	"fmt"
	"time"
)

func main() {
	go counter(1)
	time.Sleep(10 * time.Second)
}

func counter(n int) {
	i := 0
	d := time.Duration(n) * time.Second
	for {
		t := time.NewTimer(d)
		<-t.C
		time.Sleep(10 * time.Second)
		i++
		fmt.Println(i, " t ", t)
	}
}
