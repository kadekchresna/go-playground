package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calculateAreaCircle(15))
}

func calculateAreaCircle(r float32) float32 {
	return math.Phi * r * r
}
