package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// check(-1)
	switch i := randomNumber(); {
	case i > 0:
		fmt.Println("> 0")
		fallthrough
	case i < 0:
		fmt.Println("< 0")
	default:
		fmt.Println("== 0")
	}
}

func check(i int) {
	switch {
	case i > 0:
		fmt.Println("> 0")
	case i < 0:
		fmt.Println("< 0")
	default:
		fmt.Println("== 0")
	}

}

func randomNumber() int {
	return rand.Int()
}
