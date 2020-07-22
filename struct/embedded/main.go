package main

import "fmt"

type (
	person struct {
		name string
		age  int
		zip  string
		addess
	}
	addess struct {
		zip     string
		country string
	}
)

func main() {

	p := person{
		"k", 21, "123123", addess{"zip", "address"},
	}

	p2 := person{
		name: "yeeet",
		age:  23,
		zip:  "12312",
		addess: addess{
			zip:     "zip",
			country: "asdas",
		},
	}
	fmt.Println(p.zip)
	fmt.Println(p.addess.zip)
	fmt.Println(p2.zip)
	fmt.Println(p2.addess.zip)
	fmt.Println(p2)
	fmt.Printf("%v\n", p2)
	fmt.Printf("%#v\n", p2)
}
