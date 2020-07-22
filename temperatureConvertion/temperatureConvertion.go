package main

import "fmt"

func convertFahrenheitToCelcius(f float32) float32 {
	return (f - 32.0) * 5 / 9
}

func main() {
	fmt.Println(convertFahrenheitToCelcius(12.000))
}
