package main

import "fmt"

func main() {
	defer foo()
	defer bar()

}

func foo() {
	fmt.Println("FOO")
}

func bar() {
	fmt.Println("BAR")
}
