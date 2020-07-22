package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minCard = 1
	maxCard = 13
	minSuit = 1
	maxSuit = 4
)

func main() {

	rand.Seed(time.Now().UnixNano())
	//rand.Intn() will return number between inputed number -1 ex. inputed number is 12 the returned value between 0 - 11
	card := rand.Intn(maxCard-minCard) + minCard
	suit := rand.Intn(maxSuit-minSuit) + minSuit

	fmt.Println(card)
	fmt.Println(suit)

}
