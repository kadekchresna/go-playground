package main

import (
	"fmt"
	"time"
)

var delay = 100 * time.Millisecond

func main() {
	breaker := 3 * time.Second
	start := time.Now()
	load := `\|/-`
	i := 0

	for {
		if time.Since(start) > breaker {
			fmt.Println("Exiting.......")
			break

		}
		loadingAnimate(string(load[i]))
		i = (i + 1) % 4
	}

}

func loadingAnimate(symbol string) {
	fmt.Print(symbol)
	time.Sleep(delay)
	fmt.Print("\b")
}
