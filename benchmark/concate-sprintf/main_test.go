package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	fmt.Println(fmt.Sprintf("index %d", 1))
}
func BenchmarkConcat(b *testing.B) {
	i := strconv.Itoa(1)
	fmt.Println("index" + i + "")
}
