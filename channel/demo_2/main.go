package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Some Text"
	// c := make(chan string)
	b := make(chan bool)

	f := []interface{}{toLowerCase, toUpperCase}

	for i := range f {
		go f[i].(func(chan bool, string))(b, s)
	}

	// c <- s

	for range f {
		<-b
	}
	// close(c)
	close(b)
}

func toUpperCase(b chan bool, c string) {
	// s := <-c
	fmt.Println(strings.ToUpper(c))
	b <- true

}

func toLowerCase(b chan bool, c string) {
	// s := <-c
	fmt.Println(strings.ToLower(c))
	b <- true
}
