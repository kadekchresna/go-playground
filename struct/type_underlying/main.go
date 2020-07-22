package main

import (
	"fmt"
	"time"
)

type (
	person struct {
		name      string
		age       int
		lastLogin string
	}

	employee person
)

func main() {

	person1 := new(person)
	person2 := new(employee)
	person1.age = 123
	person1.lastLogin = time.Now().Format(time.RFC822)

	fmt.Printf("%T, %x, %+v\n", person1, *person1, *person1)
	fmt.Printf("%T, %x, %d", person1.age, person1.age, person1.age)
	fmt.Println()

	person2.age = 123
	fmt.Printf("%T, %x\n", person2, *person2)
	fmt.Printf("%T, %x, %d", person2.age, person2.age, person2.age)
	fmt.Println()

	// if person1 == person2 { missmatch type
	// }

	f := []interface{}{foo}

	for _, val := range f {
		fmt.Println(val.(func(int) int)(2))
	}
}

func foo(i int) int {
	return i
}
