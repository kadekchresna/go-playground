package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occured", r)
			foo()
		} else {
			fmt.Println("Application running perfectly")
		}
	}()
	panic("panic occured")
}

func foo() {
	fmt.Println("foo")
}
