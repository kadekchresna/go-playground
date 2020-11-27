package main

import "fmt"

type Human interface {
	Breathe()
}

type Joni struct {
	name string
}

func main() {
	var j Joni
	// j.name = "Joni"
	j = Joni{
		name: "Joni",
	}

	j.Breathe()

}

func (j *Joni) Breathe() {
	fmt.Println(fmt.Sprintf("%s breathing", j.name))
}
