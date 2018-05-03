package main

import (
	"fmt"
)

type Node struct {
	next, prev *Node
	payload interface{}
}

type Queue struct {
	front *Node
	rear *Node
}


// methods

func (l *Queue) CreateQueue() {
	l.front = nil
	l.rear 	= nil
}

func (l *Queue) Enqueue(payload interface{}) {
	
	node := &Node{payload: payload}
	if l.front == nil {
		l.front = node
		l.rear  = node
	} else {
		l.rear.next = node
		node.prev = l.rear
		l.rear = node
	}
}

func (l *Queue) Dequeue() {

	if l.front == nil {
		fmt.Println("Queue is empty")
	} else {
		p := l.front
		l.front = p.next
		p.next = nil
		l.front.prev = nil
	}
}

func (l *Queue) Print() {

	p := l.front
	for p != nil {
		fmt.Print("| " ,p.payload, " ")
		p = p.next
	}
	fmt.Print("|", "\n")
}

func main() {
	l := &Queue{}
	l.CreateQueue()

	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)

	l.Dequeue()
	l.Dequeue()

	l.Enqueue(8)
	l.Enqueue(9)

	l.Dequeue()

	l.Print()
}