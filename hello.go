package main

import (
	"fmt"

	"github.com/kadekchrisna/playground/black"
	"github.com/kadekchrisna/playground/design-pattern/builder/car"
)

func main() {
	fmt.Printf("olla, soi doraa~\n")
	fmt.Printf("%v\n", black.Items)

	c := car.NewCar().Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()

	c.Drive()
}
