package main

import (
	"errors"
	"fmt"
)

type Stove struct {
	brand string
}
type Frigde struct {
	brand string
}

type Furniture interface {
	Purpose() string
}

const (
	STOVE = iota
	FRIGDE
)

func main() {
	f, err := NewFurniture(FRIGDE)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(f.Purpose())
}

func NewFurniture(f int) (Furniture, error) {
	switch f {
	case STOVE:
		return new(Stove), nil
	case FRIGDE:
		return new(Frigde), nil
	default:
		return nil, errors.New("Furniture type not found")

	}
}

func (s *Stove) Purpose() string {
	return fmt.Sprintf(" i am a %T, i cook stuff", s)
}
func (f *Frigde) Purpose() string {
	return fmt.Sprintf(" i am a %T, i keep stuff cool", f)
}
