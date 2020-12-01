package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	for {
		once.Do(doThisOneTime)
	}

}

func doThisOneTime() {
	fmt.Println("doThisOneTime")
}
