package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type SingleLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func main() {
	s := Node{
		val:  2,
		next: &Node{},
	}

	fmt.Println(s)
}
