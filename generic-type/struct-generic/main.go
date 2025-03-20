package main

import "fmt"

type Person[Number int64 | float64] struct {
	Name   string
	Weight Number
}

func (p Person[int64]) PrintWeightInt() int64 {
	return p.Weight
}

func (p Person[float64]) PrintWeightFloat() float64 {
	return p.Weight
}

func main() {
	pF := new(Person[float64])
	pF.Weight = 85.2

	pI := new(Person[int64])
	pI.Weight = 100

	fmt.Println("--float--")
	fmt.Println(pF.PrintWeightFloat())
	fmt.Println(pF.PrintWeightInt())

	fmt.Println("--int--")
	fmt.Println(pI.PrintWeightFloat())
	fmt.Println(pI.PrintWeightInt())

}
