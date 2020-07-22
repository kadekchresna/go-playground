package main

import "fmt"

func main() {
	fmt.Println(nextDay(8, 6))
}
func nextDay(d int, t int) string {
	var day int = (t + d) % 7
	return fmt.Sprintf("Todah: %d, Duration: %d, Future: %d", t, d, day)
}
