package main

import (
	"fmt"
	"sync"
)

type single map[string]int

var (
	once sync.Once

	instance single
)

func New() single {
	once.Do(func() {
		instance = make(single)
	})

	return instance
}
func main() {
	s := New()

	s["a"] = 1
	fmt.Println(s["a"])
}
