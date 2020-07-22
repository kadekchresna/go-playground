package main

import "fmt"

func main() {
	// Variadic Input and Callback
	m := map[string]map[string]string{"y": {"name": "eyyy"}}
	if val := m["y"]; val == nil {
		fmt.Println("Nil")
	} else {
		fmt.Println(m["y"]["name"])
	}

	s := []int{2, 1, 3, 4}
	sum := func(n ...int) (r int) {
		for _, val := range n {
			r += val
		}
		return
	}
	fmt.Println(foo(s, sum))
}

func foo(s []int, f func(...int) int) int {
	return f(s...)
}

/* func sum(s ...int, a string)  { compile error ...int must be last parameter func sum(a string, s ...int)

} */
