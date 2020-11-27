package car

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() CommonCar
}

type CommonCar interface {
	Drive()
	Stop()
}

type car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func NewCar() Builder {
	return new(car)
}

func (c *car) Color(col Color) Builder {
	c.color = col
	return c
}

func (c *car) Wheels(w Wheels) Builder {
	c.wheels = w
	return c
}

func (c *car) TopSpeed(s Speed) Builder {
	c.speed = s
	return c
}

func (c *car) Build() CommonCar {
	return c
}

func (c *car) Drive() {
	fmt.Println("BROOMMM")
}

func (c *car) Stop() {
	fmt.Println("CITTTT")
}
