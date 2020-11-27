package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(nextDay(8, 6))
}
func nextDay(d int, t int) string {

	if 1 != 2 {
		fmt.Println(errors.New("asdas"))
	}
	var day int = (t + d) % 7
	return fmt.Sprintf("Todah: %d, Duration: %d, Future: %d", t, d, day)

}
