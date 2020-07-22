package main

import "fmt"

type (
	person struct {
		name string
		age  int
	}
)

func main() {

	person1 := person{"Yee", 23}
	person2 := person{"Yee", 23}

	fmt.Printf("%x, %x\n", &person1, &person1.age)
	fmt.Printf("%x, %x\n", &person2, &person2.age)
	compareStruct(person1, person2)

	fmt.Println(person1.age)

}

func compareStruct(sa, sb person) bool {
	fmt.Printf("%T, %x\n", sa, &sa.age)
	fmt.Printf("%T, %x\n", sb, &sb.age)
	s1 := sa
	s2 := sb
	fmt.Println("compareStruct")
	s1.age = 24
	fmt.Printf("%x, %x\n", &s1, &s1.age)
	fmt.Printf("%x, %x\n", &s2, &s2.age)
	fmt.Println(s1.age)

	if s1 == s2 {
		fmt.Println("they are the same thing")
		return true
	}

	fmt.Println("they are different")
	return false
}
