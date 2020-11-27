package main

import "fmt"

func main() {
	fmt.Println(helper("Hello"))
}

func helper(str string) string {
	if len(str) <= 1 {
		return str
	}
	return helper(str[1:]) + str[:1]
}
