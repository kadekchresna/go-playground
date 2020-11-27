package main

import "fmt"

func main() {
	var s interface{}

	s = "ass"
	switch i := s.(type) {
	case string:
		fmt.Println(fmt.Sprintf("is a %T", i))
	case int:
		fmt.Println(fmt.Sprintf("is a %T", i))
	}

}
