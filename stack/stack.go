package main

import (
	"fmt"
)

type Node struct {
	next *Node
	payload interface{}
}

type Stack struct {
	top *Node
}

func (s *Stack) CreateStack() {
	s.top = nil
}

func (s *Stack) Push(payload interface{}) {

	node := &Node {
		payload: payload,
	}

	if s.top == nil {
		s.top = node
	} else {
		node.next = s.top
		s.top = node
	}
}

func (s *Stack) Pop() {

	if s.top == nil {
		fmt.Println("Stack is empty")
	} else {
		p := s.top
		s.top = p.next
		p.next = nil
	}
}

func (s *Stack) Print() {

	p := s.top
	for p != nil {
		fmt.Println(" | ", p.payload , " | ")
		p = p.next
	}
	fmt.Println()
}

func main() {

	s := &Stack{}
	s.CreateStack()

	s.Push(5)
	s.Push(4)
	s.Push(3)
	s.Push(2)
	s.Push(1)
	s.Push(0)

	s.Print()

	s.Pop()

	s.Print()
}