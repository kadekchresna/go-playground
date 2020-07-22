package main

import "fmt"

func main() {
	fmt.Println("Proof that slice has underlying array")
	as := [...]string{"a", "b", "c"}
	s := as[0:2]
	x := 1

	fmt.Printf("s: %T, as: %T ", s, as)
	fmt.Println()
	fmt.Printf("s: %x, as: %x \n", &s, &as)
	fmt.Printf("s: %x, as: %d ", &x, x)
	fmt.Println()
	fmt.Println("len as", len(as))
	fmt.Println("cap as", cap(as))

	for i := 0; i < len(s); i++ {

		fmt.Printf("s: %x\t", &as[i])
		fmt.Printf("s: %s", as[i])
		fmt.Println()
	}
	fmt.Println("-----------------------------")

	fmt.Println("len s", len(s))
	fmt.Println("cap s", cap(s))

	for i := 0; i < len(s); i++ {

		fmt.Printf("s: %x\t", &s[i])
		fmt.Printf("s: %s", s[i])
		fmt.Println()
	}
	fmt.Println("-----------------------------")
	fmt.Printf("as: %x\t", &as[0])
	fmt.Printf("s: %x\n", &s[1])

	fmt.Printf("as: %s\t", as[0])
	fmt.Printf("s: %s", s[1])
	fmt.Println()
}
