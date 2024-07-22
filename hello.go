package main

import (
	"fmt"

	black "play.init/black/item"
	"play.init/design-pattern/builder/car"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("%v\n", black.Items)

	c := car.NewCar().Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()

	c.Drive()
}
