package main

import "testing"

func TestHighest(t *testing.T) {
	res := Highest([]int{1, 2, 3, 8, 9, 3, 2, 1})
	if res != 3 {
		t.Fatal("err")
	}
}
