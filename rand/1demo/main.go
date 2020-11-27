package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.New(rand.NewSource(2)))
}
